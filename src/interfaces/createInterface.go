package interfaces

import (
	"fmt"
	"math"
)

// Go program to Demonstrate the Interface concept

type Shape interface {
	CalcArea() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) CalcArea() float64 {
	return (t.Base * t.Height) / 2
}


func PrintType(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n",i , i)
}