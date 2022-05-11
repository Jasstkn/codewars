package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func main() {
	fmt.Println(solution("banana", "ana"))
}
