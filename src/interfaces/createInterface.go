package interfaces

import (
	"fmt"
	"math"
)

// Go program to Demonstrate the Interface concept

type Shape interface {
	CalcArea() float64
}

// create Shape Circle
type Circle struct {
	Radius float64
}

// calculates Shape Circle
func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// create Shape Triangle
type Triangle struct {
	Base float64
	Height float64
}

// calculates Area of Triangle
func (t Triangle) CalcArea() float64 {
	return (t.Base * t.Height) / 2
}


// create Shape Rectange
type Rectangle struct {
	Length float64
	Breadth float64
}

// calculates Area of Rectangle
func (r Rectangle) CalcArea() float64 {
	return r.Length * r.Breadth
}

func PrintType(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n",i , i)
}

func main(){
	var shape Shape

	shape = Circle{Radius: 7.0}
	fmt.Println("Area of Circle: ", shape.CalcArea())

	shape = Triangle{Base: 5, Height: 6}
	fmt.Println("Area of Triangle: ", shape.CalcArea())

	// Interface to Print the dynamic Type initialization
	PrintType(10)
	PrintType("10")
	PrintType(10.932872)
	PrintType(10/2==0)
	PrintType([]int{})
	PrintType(map[int]int{})
}