package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func ToCsvText(array [][]int) string {
	rows := make([]string, 0, len(array))
	for _, row := range array {
		rows = append(rows, strings.Join(SliceIntToAscii(row), ","))
	}

	return strings.Join(rows, "\n")
}

func SliceIntToAscii(array []int) []string {
	arrayStr := make([]string, 0, len(array))
	for _, n := range array {
		arrayStr = append(arrayStr, strconv.Itoa(n))
	}
	return arrayStr
}

func main() {
	start := time.Now()
	fmt.Println(ToCsvText([][]int{
		{-25, 21, 2, -33, 48},
		{30, 31, -32, 33, -34}}))
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
