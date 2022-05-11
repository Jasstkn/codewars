package main

import (
	"fmt"
	"sort"
)

func Gps(s int, x []float64) int {
	if len(x) <= 1 {
		return 0
	}
	avgSpeed := make([]float64, len(x)-1)
	for i := 0; i < len(x)-1; i++ {
		avgSpeed[i] = (3600 * (x[i+1] - x[i])) / float64(s)
	}
	sort.Float64s(avgSpeed)
	return int(avgSpeed[len(avgSpeed)-1])
}

func main() {
	fmt.Println(Gps(20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}))
}
