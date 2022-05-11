package main

import (
	"fmt"
	"strings"
)

func MultiTable(number int) string {
	var b strings.Builder
	for i := 1; i < 11; i++ {
		b.WriteString(fmt.Sprintf("%d * %d = %d", i, number, i * number))
		if i != 10 {
			b.WriteString("\n")
		}

	}
	return b.String()
}

func main() {
	fmt.Print(MultiTable(5))
}
