package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	sort.Ints(numbers)
	return numbers
}

func main() {
	fmt.Print(SortNumbers([]int{}))
}
