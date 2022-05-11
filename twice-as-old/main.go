package main

import "fmt"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	if dadYearsOld < sonYearsOld*2 {
		return sonYearsOld*2 - dadYearsOld
	} else {
		return dadYearsOld - sonYearsOld*2
	}

	// return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
}

func main() {
	fmt.Print(TwiceAsOld(36, 7))
}
