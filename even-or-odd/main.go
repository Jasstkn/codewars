package main

import (
	"fmt"
	"log"
	"time"
)

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	start := time.Now()
	fmt.Println(EvenOrOdd(1))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
