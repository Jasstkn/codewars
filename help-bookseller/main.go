package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	m := make(map[string]int)
	for _, v := range listCat {
		m[v] = 0
	}

	for _, v := range listArt {
		_, isExist := m[v[:1]]
		if isExist {
			val, _ := strconv.Atoi(strings.Split(v, " ")[1])
			m[v[:1]] += val
		}
	}

	r := make([]string, len(m))

	for i, v := range listCat {
		r[i] = fmt.Sprintf("(%s : %d)", v, m[v])
	}

	return strings.Join(r, " - ")
}

func main() {
	fmt.Println(StockList(
		[]string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
		[]string{"A", "B", "C", "D"}))
}
