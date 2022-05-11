package main

import (
	"fmt"
)

func checkParams(h, bounce, window float64) bool {
	if h <= 0 {
		return false
	}
	if bounce <= 0 || bounce >= 1 {
		return false
	}
	if window >= h {
		return false
	}
	return true
}

func BouncingBall(h, bounce, window float64) int {
	isValid := checkParams(h, bounce, window)
	if !isValid {
		return -1
	}
	counter := 1
	h = h * bounce
	for h > window {
		h *= bounce
		counter += 2
	}
	return counter
}

func main() {
	fmt.Println(BouncingBall(40, 0.4, 10))
}
