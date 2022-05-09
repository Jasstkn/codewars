package main

import "fmt"

func ReverseSeq(n int) []int {

	var a []int

	a = make([]int, n)

	for i := n; i > 0; i-- {
		a[n-i] = i
	}
	return a
}

func main() {
	fmt.Println(ReverseSeq(5))
}
