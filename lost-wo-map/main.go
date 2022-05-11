package main

import (
	"fmt"
	"log"
	"time"
)

func Maps(x []int) (res []int) {
	res = make([]int, len(x))
	for i := 0; i < len(x); i++ {
		res[i] = x[i] * 2
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(Maps([]int{1, 2, 3}))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
