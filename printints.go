package main

import (
    "fmt"
    "bytes"
)

func intsToString(value []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range value {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

func main()  {
    fmt.Println(intsToString([]int{1,2,3}))
    fmt.Println([]int{2,3,4})
}
