// https://www.codewars.com/kata/5d59576768ba810001f1f8d6/go

package main

func Quadratic(x1, x2 int) [3]int {
	return [3]int{1, -x1 + -x2, x1 * x2}
}

func main() {
	for _, i := range Quadratic(4, 9) {
		print(i, " ")
	}
}
