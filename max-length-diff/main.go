package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	var maxDiff float64

	for _, v1 := range a1 {
		for _, v2 := range a2 {
			diff := math.Abs(float64(len(v1) - len(v2)))
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	return int(maxDiff)
}

func main() {
	start := time.Now()
	fmt.Println(MxDifLg([]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}, []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
