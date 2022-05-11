package main

import (
	"fmt"
	"math"
	"strings"
)

func GetListOfChunks(s string, n int) []string {
	chunkNum := len(s) / n
	l := make([]string, chunkNum)
	for i := 0; i < chunkNum; i++ {
		l[i] = s[i*n : n*i+n]
	}

	return l
}

func isReversible(chunk string) bool {
	var sum int
	for _, c := range chunk {
		sum += int(math.Pow(float64(c), 3))
	}
	if sum%2 == 0 {
		return true
	}
	return false
}

func reverse(chunk string) string {
	runes := []rune(chunk)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func rotate(chunk string) string {
	return chunk[1:] + string(chunk[0])
}

func Revrot(s string, n int) string {
	if len(s) == 0 || n <= 0 || len(s) < n {
		return ""
	}
	chunksList := GetListOfChunks(s, n)

	for i, c := range chunksList {
		if isReversible(c) {
			chunksList[i] = reverse(c)
		} else {
			chunksList[i] = rotate(c)
		}
	}

	return strings.Join(chunksList, "")
}

func main() {
	fmt.Println(Revrot("733049910872815764", 5))
}
