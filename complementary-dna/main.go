package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func DNAStrand(dna string) string {
	// for i, v := range dna {
	// 	if string(v) == "A" {
	// 		dna = dna[:i] + string("T") + dna[i+1:]
	// 	}
	// 	if string(v) == "T" {
	// 		dna = dna[:i] + string("A") + dna[i+1:]
	// 	}
	// 	if string(v) == "C" {
	// 		dna = dna[:i] + string("G") + dna[i+1:]
	// 	}
	// 	if string(v) == "G" {
	// 		dna = dna[:i] + string("C") + dna[i+1:]
	// 	}
	// switch string(v) {
	// case "A":
	// 	dna = dna[:i] + string("T") + dna[i+1:]
	// case "T":
	// 	dna = dna[:i] + string("A") + dna[i+1:]
	// case "C":
	// 	dna = dna[:i] + string("G") + dna[i+1:]
	// case "G":
	// 	dna = dna[:i] + string("C") + dna[i+1:]
	// }
	// }

	var dnaReplacer *strings.Replacer = strings.NewReplacer(
		"A", "T",
		"T", "A",
		"C", "G",
		"G", "C",
	)

	return dnaReplacer.Replace(dna)
}

func main() {
	start := time.Now()

	fmt.Println(DNAStrand("AAAA"))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
