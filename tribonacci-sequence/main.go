package main

import (
	"fmt"
	"log"
	"time"
)

func Tribonacci(signature [3]float64, n int) []float64 {
	// fail fast when expected number of elements == 0
	if n == 0 {
		return []float64{}
	}

	// initialize slice
	res := make([]float64, n)

	// fill the elements
	for i := 0; i < len(res); i++ {
		if i < len(signature) {
			res[i] = signature[i]
		} else {
			res[i] = res[i-1] + res[i-2] + res[i-3]
		}
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 1))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
