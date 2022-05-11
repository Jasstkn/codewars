package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func combat(health, damage float64) float64 {
	return math.Max(health-damage, 0)
}

func main() {
	start := time.Now()
	fmt.Println(combat(1.0, 100))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
