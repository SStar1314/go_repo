package main

import (
    "fmt"
    "os"
)

func main()  {
    fmt.Println(os.Getwd())
    fmt.Println(os.Hostname())
    fmt.Println(os.Getpid())
    fmt.Println(os.Getppid())
    fmt.Println(os.Getuid())
}
