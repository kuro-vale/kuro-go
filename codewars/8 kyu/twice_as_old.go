// https://www.codewars.com/kata/5b853229cfde412a470000d0/go

package main

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	return int(math.Abs(float64(dadYearsOld - sonYearsOld*2)))
}

func main() {
	print(TwiceAsOld(36, 7))
}
