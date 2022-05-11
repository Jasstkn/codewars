package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func FindOdd(seq []int) (res int) {
	sort.Ints(seq)

	counterMap := make(map[int]int)
	for _, v := range seq {
		if _, ok := counterMap[v]; ok {
			counterMap[v]++
		} else {
			counterMap[v] = 1
		}
	}

	for i, v := range counterMap {
		if v%2 != 0 {
			res = i
		}
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
