package main

import (
	"fmt"
	"strings"
)

func VertMirror(s string) string {
	a := strings.Split(s, "\n")
	for p, v := range a {
		runes := []rune(v)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		a[p] = string(runes)
	}
	return strings.Join(a, "\n")
}
func HorMirror(s string) string {
	a := strings.Split(s, "\n")
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	return strings.Join(a, "\n")
}

type FParam func(string) string
func Oper(f FParam, x string) string {
	return f(x)
}

func main() {
	fmt.Printf(Oper(VertMirror, "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"))
}
