package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countSheep(num int) string {
	const phrase = " sheep..."
	var buffer strings.Builder
	for i := 1; i < num+1; i++ {
		buffer.WriteString(strconv.Itoa(i))
		buffer.WriteString(phrase)
	}

	return buffer.String()

	// for count := 1; count <= num; count++ {
	//     fmt.Fprintf(&sb, "%d sheep...", count)
	// }

	// return sb.String()
}

func main() {
	fmt.Println(countSheep(5))
}
