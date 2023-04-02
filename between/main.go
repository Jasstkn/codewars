package main

import (
	"fmt"
	"log"
	"time"
)

func Between(a, b int) []int {
	nums := make([]int, 0)
	for i := a; i <= b; i++ {
		nums = append(nums, i)
	}
	return nums
}

func main() {
	start := time.Now()
	fmt.Println(Between(1, 4))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
