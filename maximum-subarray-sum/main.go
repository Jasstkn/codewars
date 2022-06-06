package main

import (
	"fmt"
	"log"
	"time"
)

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	currentSum, previousSum := 0, 0
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			currentSum += numbers[j]
			if currentSum >= previousSum {
				previousSum = currentSum
			}
		}
		currentSum = 0
	}

	return previousSum
}

func main() {
	start := time.Now()
	fmt.Println(MaximumSubarraySum([]int{}))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
