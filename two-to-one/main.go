package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func TwoToOne(s1 string, s2 string) (res string) {
	s := s1 + s2
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s, char) {
			res += char
		}
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
