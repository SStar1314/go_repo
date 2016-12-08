package main

import (
    "fmt"
    "time"
)

func main() {
    naturals := make(chan int)
    squares := make(chan int)

    go func ()  {
        for x := 0; ; x++ {
            naturals <- x
            time.Sleep(1000*time.Millisecond)
        }
    }()
    go func ()  {
        for {
            x := <- naturals
            squares <- x * x
        }
    }()

    for {
        fmt.Println(<- squares)
    }
}
