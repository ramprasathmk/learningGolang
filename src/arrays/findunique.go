package arrays

// function to check the array has unique elements or not using map concept
func CheckUnique(array []int) bool {
	cache := make(map[int]bool) 
	
	for _, elem := range array {
		
		if cache[elem] {
			return false
		}
		
		cache[elem] = true
	}
	return true
}

// function to check the array has unique elements or not using map with struct concept
func CheckUnique2(array []int) bool {
	cache := make(map[int]struct{})

	for _, elem := range array {
		cache[elem] = struct{}{}
	}

	return len(array) == len(cache)
}

// add non-unique elements to a slice
func AddUnique(array []int) (bool, []int, []int) {
	
	// unique elements array
	uArray := make([]int, 0)
	lArray := make([]int, 0)

	// create a map 
	cache := make(map[int]bool)

	for _, elem := range array {
		if cache[elem] {
			uArray = append(uArray, elem)
			return false, uArray, nil
		} else {
			lArray = append(lArray, elem)
			return true, lArray, nil
		}
	}
	return true, lArray, uArray
}

func PrintUArr(arr []int) {}