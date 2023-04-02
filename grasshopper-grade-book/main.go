package main

import (
	"fmt"
	"log"
	"time"
)

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3
	switch {
	case avg < 60:
		return 'F'
	case avg < 70:
		return 'D'
	case avg < 80:
		return 'C'
	case avg < 90:
		return 'B'
	default:
		return 'A'
	}
}

func main() {
	start := time.Now()
	fmt.Println(GetGrade(100, 100, 100))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
