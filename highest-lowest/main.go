package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func HighAndLow(in string) string {
	inList := strings.Split(in, " ")
	res := make([]int, 0, len(inList))

	for _, v := range inList {
		el, _ := strconv.Atoi(v)
		res = append(res, el)
	}
	sort.Ints(res)
	return fmt.Sprintf("%d %d", res[len(res)-1], res[0])
}

func main() {
	start := time.Now()
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
