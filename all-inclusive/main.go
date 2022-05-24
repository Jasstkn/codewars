package main

import (
	"fmt"
	"log"
	"time"
)

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}

	combinations := make([]string, len(strng))

	for i := 0; i < len(combinations); i++ {
		strng = string(strng[1:]) + string(strng[0])
		combinations[i] = strng
	}

	var counter int
	for _, v1 := range combinations {
		for _, v2 := range arr {
			if v1 == v2 {
				counter++
				break
			}
		}
	}

	if counter == len(strng) {
		return true
	}

	return false
}

func main() {
	start := time.Now()
	fmt.Println(ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
