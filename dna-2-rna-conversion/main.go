package main

import (
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string {
	var buffer strings.Builder
	for _, v := range dna {
		if string(v) == "T" {
			buffer.WriteString("U")
		} else {
			buffer.WriteString(string(v))
		}
	}
	return buffer.String()
}

func DNAtoRNAEff(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}
