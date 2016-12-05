package main

import (
    "fmt"
)

func main()  {
    var a [3]int = [3]int{1, 2, 3}
    for i, v := range a {
        fmt.Printf("%d %d\n", i, v)
    }
}
