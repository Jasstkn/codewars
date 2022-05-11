package main

import (
	"fmt"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}

	for i, v := range array1 {
		array1[i] = v * v
	}
	sort.Ints(array1)
	sort.Ints(array2)
	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}))
}
