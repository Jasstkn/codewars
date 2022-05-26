package main

import (
	"fmt"
	"log"
	"time"
)

func HowMuch(m int, n int) [][3]string {
	var res [][3]string
	if n < m {
		m, n = n, m
	}
	for i := m; i < n+1; i++ {
		c := i / 9
		b := i / 7
		if i-c*9 == 1 && i-b*7 == 2 {
			res = append(res, [3]string{
				fmt.Sprintf("M: %d", i),
				fmt.Sprintf("B: %d", b),
				fmt.Sprintf("C: %d", c)})
		}
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(HowMuch(10000, 9950))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
