package main

import (
	"fmt"
	"log"
	"time"
	"unicode"
)

func Capitalize(st string, arr []int) string {
	runes := []rune(st)
	for _, v := range arr {
		if v < len(st) {
			runes[v] = unicode.ToUpper(runes[v])
		}
	}

	return string(runes)
}

func main() {
	start := time.Now()
	fmt.Println(Capitalize("abcdef", []int{1, 2, 5, 100}))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
