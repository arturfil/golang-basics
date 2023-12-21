package interfaces

import (
	"fmt"
	"math"
	"reflect"
)

// model behavior

// interfaces in Go are implicitly implemented

type Square struct {
    NumberOfSides int 
    LengthOfSide float32
}

func (s *Square) GetArea() float32 {
   return float32(s.NumberOfSides) * s.LengthOfSide 
}

type Triangle struct {
    NumberOfSides int
    Base float32
    Height float32
}

func (t *Triangle) GetArea() float32 {
    return (t.Base * t.Height) / 2
}

type Circle struct {
    radius float32
}

func (c *Circle) GetArea() float32 {
    return 3.1416 * float32(math.Pow(2, float64(c.radius)))
}

type AreaGetter interface {
    GetArea() float32
}

func AreaCalculator(a AreaGetter) float32 {
    fmt.Println(reflect.TypeOf(a).Elem(), "has an area of:", a.GetArea())
    return a.GetArea()
}

func InterfacesMain() {

    s := &Square{NumberOfSides: 4, LengthOfSide: 2.4}
    t := &Triangle{NumberOfSides: 3, Base: 2, Height: 3.3}
    s2 := &Square{NumberOfSides: 4, LengthOfSide: 3.2}
    c := &Circle{radius: 2.4}

    shapes := []AreaGetter {s, t, s2, c}

    for _, shape := range shapes {
        AreaCalculator(shape)
    }
    
}
