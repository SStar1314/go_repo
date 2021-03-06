package main

import (
    "math"
    "fmt"
)

type Point struct {
    X, Y float64
}

func Distance(p, q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64)  {
    p.X *= factor
    p.Y *= factor
}

func main()  {
    r := &Point{1,2}
    r.ScaleBy(2)
    fmt.Println(*r)

    p := Point{1, 2}
    q := Point{4, 6}

    fmt.Println(Distance(p, q))
    fmt.Println(p.Distance(q))

    perim := Path {
        {1,1}, {5,1}, {5,4}, {1,1},
      }
    fmt.Println(perim.Distance())
}

type Path []Point

func (path Path) Distance() float64 {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
        }
    }
    return sum
}
