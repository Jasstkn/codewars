package main

import (
	"fmt"
	"log"
	"time"
	"unicode"
)

func Capitalize(st string) []string {
	even, odd := []rune(st), []rune(st)
	for i, v := range st {
		if i%2 == 0 {
			even[i] = unicode.ToUpper(v)
		} else {
			odd[i] = unicode.ToUpper(v)
		}
	}
	return []string{string(even), string(odd)}
}

func main() {
	start := time.Now()

	fmt.Println(Capitalize("codewars"))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
