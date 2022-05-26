package main

import (
	"fmt"
	"log"
	"time"
)

func Xor(a, b bool) bool {
	return a != b
}

func main() {
	start := time.Now()
	fmt.Println(Xor(false, Xor(false, false)))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
