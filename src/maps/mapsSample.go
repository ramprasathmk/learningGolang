package maps

import "fmt"

func RunMaps() {
	fmt.Println("Maps Demo")
	// map creation
	emptyMap := make(map[int]int)
	fmt.Println("Empty map: ", emptyMap)

	// Fruits Map
	var fruitsMap map[string]float64
	fruitsMap = map[string]float64{"Mango": 50.0, }
	fmt.Println(fruitsMap)
}
