package main

import (
	"fmt"
	"sort"
)

func ExpressionMatter(a int, b int, c int) int {
	calculatedNumbers := []int{
		a + b + c,
		a * b * c,
		a*b + c,
		a + b*c,
		(a + b) * c,
		a * (b + c),
	}

	sort.Ints(calculatedNumbers)
	return calculatedNumbers[len(calculatedNumbers)-1]
}

func main() {
	fmt.Println(ExpressionMatter(2, 1, 1))
}
