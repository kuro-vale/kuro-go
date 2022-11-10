// https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097/go
package main

import (
	"fmt"
	"math"
)

func century(year int) int {
	return int(math.Ceil(float64(year) / 100))
}

func main() {
	fmt.Printf("%d", century(2001))
}
