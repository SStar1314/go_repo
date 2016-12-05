package main

import (
    "fmt"
    "os"
    "strconv"
)

func main()  {
    if len(os.Args) < 2 {
        fmt.Println("Error arguments number")
    }
    n, _ := strconv.Atoi(os.Args[1])

    fibonacci(n)
}

func fibonacci(n int)  {
    var x, y int64 = 0, 1
    for i := 0; i < n; i ++ {
        fmt.Printf("n %d: %d \n", i, x)
        x, y = y, x+y
    }
}
