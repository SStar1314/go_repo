package main

import (
    "flag"
    "fmt"
    "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main()  {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
    f()
    fmt.Println(*global)
    fmt.Println(global)
}

func delta(old, new int) int {
    return  new - old
}

var global *int

func f()  {
    var x int
    x = 1
    global = &x
    fmt.Println(&x)
}
