package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) (res string) {
	for _, c := range strings.ToLower(word) {
		if strings.Count(strings.ToLower(word), string(c)) > 1 {
			res += ")"
		} else {
			res += "("
		}
	}
	return res
}

func main() {
	fmt.Println(DuplicateEncode("Success"))
}
