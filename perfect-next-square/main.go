package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func FindNextSquare(sq int64) int64 {
	sqrt := int64(math.Sqrt(float64(sq)))
	if sq%sqrt != 0 {
		return -1
	}
	sqrt++
	return sqrt * sqrt
}

func main() {
	start := time.Now()
	fmt.Println(FindNextSquare(121))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
