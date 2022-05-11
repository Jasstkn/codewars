package main

import (
	"fmt"
	"strings"
)

func Repeater(s string, n int) string {
	return strings.Repeat(s, n)
}

func main() {
	fmt.Println(Repeater("a", 5))
}
