package main

import (
    "fmt"
    "os"
    "strconv"
)

func main()  {
    var x, y int
    if len(os.Args) < 3 {
        fmt.Println("Error arguments number")
    }
    x, _ = strconv.Atoi(os.Args[1])
    y, _ = strconv.Atoi(os.Args[2])

    fmt.Printf("%d, %d\n", x, y)

    c := gcd(x, y)

    fmt.Printf("%d, %d greatest common divisor is %d\n", x, y, c)
}

func gcd(x,y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}
