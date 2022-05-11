package main

func NbYear(p0 int, percent float64, aug int, p int) int {

	var y int
	var population float64

	for p0 < p {
		population = float64(p0) + float64(p0)*(percent/100) + float64(aug)
		p0 = int(population)
		y++
	}

	return y

	// if p0 >= p {
	// 	return n
	// }
	//   newP := p0 + aug + int(float64(p0)*percent/100)
	//   n++
	// return n + NbYear(newP, percent, aug, p)
}

func main() {
	NbYear(1500, 5, 100, 5000)
}
