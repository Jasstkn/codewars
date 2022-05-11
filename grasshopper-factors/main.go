package main

import "fmt"

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}

func main() {
	fmt.Println(CheckForFactor(10, 2))
}
