package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	n := len(str)
	for i := 0; i < n; i++ {
		n -= 1
		if str[i] != str[n] {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	fmt.Println(IsPalindrome("Abba"))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
