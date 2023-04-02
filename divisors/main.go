package main

import (
	"fmt"
	"log"
	"time"
)

func Divisors(n int) int {
	divs := 1 // always can be divided by itself

	for i := 1; i < n/2; i++ {
		if n%i == 0 {
			divs++
		}
	}
	return divs
}

func main() {
	start := time.Now()
	fmt.Println(Divisors(1))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
