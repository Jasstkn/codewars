package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func solve(slice []string) (res []int) {
	for _, word := range slice {
		counter := 0
		for i, c := range strings.ToUpper(word) {
			if string(c) == string(rune('A'+i)) {
				counter++
			}
		}
		res = append(res, counter)
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(solve([]string{"abode", "ABc", "xyzD"}))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
