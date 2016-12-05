package main

import (
    "fmt"
)

type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

type Flags uint

const (
    FlagUp Flags = 1 << iota
    FlagBroadcast
    FlagLoopback
    FlagPointToPoint
    FlagMulticast
)

func main()  {
    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
    fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}
