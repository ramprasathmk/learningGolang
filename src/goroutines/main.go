package main

import (
	"fmt"
	"log"
	"time"
)

// someFunc function prints the number in string type
func someFunc(number string) {
	fmt.Println("goroutine: " + number)
}

// loopFunc function loops upto n times
func loopFunc(n int) {
	fmt.Printf("loop running %d times\n", n)
	i := 0
	for range n {
		i++
	}
}

func main() {
	// initial log message
	log.Println("Program Start...")

	go someFunc("1")
	go someFunc("2")
	go loopFunc(3)
	go loopFunc(4)
	go loopFunc(5)

	time.Sleep(time.Second * 1)

	fmt.Println("Sample Message: Hello")
	
	// final log message
	log.Println("Program End...")
}