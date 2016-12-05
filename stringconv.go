package main

import (
    "fmt"
    "strconv"
)

func main()  {
    x := 123
    y := fmt.Sprintf("%d", x)
    fmt.Println(y, strconv.Itoa(x))
    fmt.Println(strconv.FormatInt(int64(x), 2))
    fmt.Println(strconv.FormatInt(int64(x), 4))
    fmt.Println(strconv.ParseInt(y, 10, 64))
}
