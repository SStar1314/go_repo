package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
    //"text/template"
    "html/template"
)

const IssuesURL = "https://api.github.com/search/issues"
const temp = `{{.TotalCount}} issues:
    {{range .Items}}------------------------------------
    Number: {{.Number}}
    User:   {{.User.Login}}
    Title:  {{.Title | printf "%.64s"}}
    Age:    {{.CreatedAt | daysAgo}} days
    {{end}}`

type IssuesSearchResult struct {
    TotalCount    int
    Items         []*Issue
}

type Issue struct {
    Number      int
    HTMLURL     string
    Title       string
    State       string
    User        *User
    CreatedAt   time.Time
    Body        string
}

type User struct {
    Login       string
    HTMLURL     string
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(IssuesURL + "?q=" + q)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }
    resp.Body.Close()
    return &result, nil
}

func daysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}

//var report = template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp))

var issueList = template.Must(template.New("issuelist").Parse(`
    <h1>{{.TotalCount}} issues</h1>
     <table>
     <tr style='text-align: left'>
       <th>#</th>
       <th>State</th>
       <th>User</th>
       <th>Title</th>
     </tr>
     {{range .Items}}
     <tr>
       <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
       <td>{{.State}}</td>
       <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
       <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
     </tr>
     {{end}}
     </table>
  `))

var autotmpl = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

func main()  {
    result, err := SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
    }
    fmt.Println("-----------------------------------")
//    if err := issueList.Execute(os.Stdout, result); err != nil {
//        log.Fatal(err)
//    }

    t := template.Must(template.New("escape").Parse(autotmpl))
    var data struct {
        A     string
        B     template.HTML
    }
    data.A = "<b>Hello!</b>"
    data.B = "<b>Hello!</b>"
    if err := t.Execute(os.Stdout, data); err != nil {
        log.Fatal(err)
    }
}
