package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func AbbrevName(name string) string {
	words := strings.Split(strings.ToUpper(name), " ")

	return words[0][0:1] + "." + words[1][0:1]
}

func main() {
	start := time.Now()

	fmt.Println(AbbrevName("Sam Harris"))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
