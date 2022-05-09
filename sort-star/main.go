package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoSort(arr []string) string {

	var buffer strings.Builder

	sort.Strings(arr)
	fmt.Println(arr[0])
	for i, v := range arr[0] {
		buffer.WriteString(string(v))
		if i != len(arr[0])-1 {
			buffer.WriteString("***")
		}
	}
	return buffer.String()

	// chars := strings.Split(arr[0], "")
	// return strings.Join(chars, "***")
}

func main() {
	fmt.Println(TwoSort([]string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}))
}
