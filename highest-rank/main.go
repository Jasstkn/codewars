package main

import (
	"fmt"
	"log"
	"time"
)

func HighestRank(nums []int) int {
	var mode, number int
	modeMap := make(map[int]int, len(nums))
	for _, n := range nums {
		modeMap[n] += 1
	}
	for n, m := range modeMap {
		if m > mode {
			mode = m
			number = n
		}
		if m == mode && n > number {
			number = n
		}
	}
	return number
}

func main() {
	start := time.Now()
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
