package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

func IterPi(epsilon float64) (int, string) {
	var leibnizPi float64
	var iter int

	denominator := float64(1)

	for i := 0; isConverged(leibnizPi, epsilon); i++ {
		if i%2 == 0 {
			leibnizPi += 4 / denominator
		} else {
			leibnizPi -= 4 / denominator
		}
		denominator += 2
		iter++
	}
	return iter, strconv.FormatFloat(leibnizPi, 'f', 10, 64)
}

func isConverged(value float64, epsilon float64) bool {
	if math.Abs(value-math.Pi) <= epsilon {
		return false
	}
	return true
}

func main() {
	start := time.Now()
	fmt.Println(IterPi(0.01))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
