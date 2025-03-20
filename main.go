package main

import (
	// "github.com/ramprasathmk/learningGolang/src"
	// "github.com/ramprasathmk/learningGolang/src/maps"
	// "github.com/ramprasathmk/learningGolang/src/structs"
	"fmt"

	"github.com/ramprasathmk/learningGolang/src/interfaces"
)

func main() {
	// src.Hi()
	// structs.RunStructure()
	// maps.RunMaps()

	var shape interfaces.Shape

	shape = interfaces.Circle{Radius: 7.0}
	fmt.Println("Area of Circle: ", shape.CalcArea())

	shape = interfaces.Triangle{Base: 5, Height: 6}
	fmt.Println("Area of Triangle: ", shape.CalcArea())

	// Interface to Print the dynamic Type initialization
	interfaces.PrintType(10)
	interfaces.PrintType("10")
	interfaces.PrintType(10.932872)
	interfaces.PrintType(10/2==0)
	interfaces.PrintType([]int{})
	interfaces.PrintType(map[int]int{})
}
