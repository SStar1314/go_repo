package main

import (
    "fmt"
    "sort"
)

func main()  {
    ages := make(map[string]int)
    ages["jack"] = 21
    var names []string
    for name, _ := range ages {
        names = append(names, name)
    }
    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
}
