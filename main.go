package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/ramprasathmk/learningGolang/src/arrays"
	API "github.com/ramprasathmk/learningGolang/src/API"
)

func main() {
	fmt.Println("Starting main() function")
	// new array
	// array := []int{1,2,3, 5, 6, 2, 4, 2, 1,4 ,6, 4}
	// fmt.Println(array, arrays.CheckUnique(array))
	// array2 := []int{1, 2, 3, 5, 6}
	// fmt.Println(array2, arrays.CheckUnique(array2))
	// log.Println("Array 1: ", array)
	// lcheck, larr, uarr := arrays.AddUnique(array)
	// fmt.Printf("check: %v, larr: %v, uarr: %v\n", lcheck, larr, uarr)
	log.Println("Server Started...")
	// entrypoints
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "Hello, World from the Go Server")
	})
	http.HandleFunc("/hi", API.Greet)
	http.HandleFunc("/getuserdata", API.GetUserData)
	http.HandleFunc("/getuserlist", API.GetUserList)
	// port
	http.ListenAndServe(":29100", nil)
	log.Println("Server Ended...")
}
