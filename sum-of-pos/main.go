package main

import "fmt"

func PositiveSum(numbers []int) (sum int) {
	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}

func main() {
	fmt.Print(PositiveSum([]int{1, -2, 3, 4, 5}))
}
