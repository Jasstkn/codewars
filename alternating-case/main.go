package main

import (
	"fmt"
	"strings"
)

func ToAlternatingCase(str string) string {
	var buffer strings.Builder

	for _, v := range str {
		if strings.ToLower(string(v)) == string(v) {
			buffer.WriteString(strings.ToUpper(string(v)))
		} else {
			buffer.WriteString(strings.ToLower(string(v)))
		}

	}

	return buffer.String()

	// alternate := func(r rune) rune {
	//     if unicode.IsLower(r) {
	//       return unicode.ToUpper(r)
	//     } else {
	//       return unicode.ToLower(r)
	//     }
	//   }

	// return strings.Map(alternate, str)
}

func main() {
	fmt.Println(ToAlternatingCase("hello world"))
}
