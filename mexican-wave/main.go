package main

import (
	"fmt"
	"strings"
)

func wave(words string) (res []string) {
	// Your code here and happy coding!
	res = []string{}
	for i, v := range words {
		if string(v) != " " {
			res = append(res, words[:i]+strings.ToUpper(string(v))+words[i+1:])
		}
	}

	return res
}

func main() {
	fmt.Println(wave(" x yz"))
}
