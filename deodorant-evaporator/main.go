package main

import "fmt"

func Evaporator(content float64, evapPerDay int, threshold int) int {
	var d int

	contentS := content

	for content >= contentS*(float64(threshold)/100) {
		d++
		content -= content * (float64(evapPerDay) / 100)
	}

	return d
}

func main() {
	fmt.Println(Evaporator(10, 10, 10))
}
