package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func Solution(list []int) (result string) {
	result += strconv.Itoa(list[0])
	for i := 1; i < len(list)-1; i++ {
		if list[i]-1 == list[i-1] && list[i]+1 == list[i+1] {
			result += strconv.Itoa(list[i])
		}
	}

	return result
}

func main() {
	start := time.Now()
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
