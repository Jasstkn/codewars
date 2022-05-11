package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, c := range vowels {
		count += strings.Count(str, c)
	}
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
