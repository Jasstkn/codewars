package main

import (
	"fmt"
	"strings"
)

func WordsToMarks(s string) int {
	var wordValue, counter int
	for ch := 'a'; ch <= 'z'; ch++ {
		counter++
		if strings.Count(s, string(ch)) != 0 {
			wordValue += counter * strings.Count(s, string(ch))
		}
	}
	return wordValue

	// t := map[rune]int{}
    // for i, c := range "abcdefghijklmnopqrstuvwxyz" { t[c] = i + 1 }

    // res := 0
    // for _, c := range s { res += t[c] }
}

func main() {
	fmt.Println(WordsToMarks("attitude"))
}
