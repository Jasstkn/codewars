package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	var counter int

	for i := 0; i < n+1; i++ {
		counter += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return counter
}

func main() {
	fmt.Println(NbDig(550, 5))
}
