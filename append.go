package main

import (
    "fmt"
)

func main()  {
    var x, y []int
    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
    y = appendInt2(x, 11,12,13)
    fmt.Printf("%d cap=%d\t%v\n", -1, cap(y), y)
}

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        z = x[:zlen]
    } else {
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2*len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[len(x)] = y
    return z
}

func appendInt2(x []int, y ...int) []int {
    var z []int
    zlen := len(x) + len(y)
    z = make([]int, zlen, zlen+1)
    copy(z, x)
    copy(z[len(x):], y)
    return z
}
