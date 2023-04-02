package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func TowerBuilder(nFloors int) []string {
	tower := make([]string, 0)

	for i := 0; i < nFloors; i++ {
		spaces := strings.Repeat(" ", nFloors-(i+1))
		tower = append(tower, spaces+strings.Repeat("*", 2*i+1)+spaces)
	}

	return tower
}

func main() {
	start := time.Now()
	fmt.Println(TowerBuilder(3))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
