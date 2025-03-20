package maps

import "fmt"

func CreateMaps() {
	fmt.Println("Creating Maps")
	// map creation using make()
	emptyMap := make(map[int]int)
	fmt.Println("Empty map: ", emptyMap)

	// `Fruits Map`: map declaration without initialization
	var fruitsMap map[string]float64
	fruitsMap = map[string]float64{"Mango": 50.0, }
	fmt.Println("Fruits Map", fruitsMap)

	// map declaration with initializing values using `:=` shorthand operator
	// Key: Planets Name, Value: Number of Moons
	MoonCountOfPlanetsMap := map[string]int{
		"Earth": 1, 
		"Mars": 2, 
		"Jupiter": 95, 
		"Saturn": 146, 
		"Uranus": 28, 
		"Neptune": 16,
		"pluto": 5, // Dwarf Planet
	}
	fmt.Println("Number of Moons Map: ", MoonCountOfPlanetsMap)

	for Planet, Moons := range MoonCountOfPlanetsMap {
		fmt.Printf("Planet: %v, Moons: %v\n", Planet, Moons)
	} 

	calcMoons := 0
	for _, moonCount := range MoonCountOfPlanetsMap {
		calcMoons += moonCount
	}

	fmt.Println("Current tally of moons orbiting planets: ", calcMoons)
}