// https://www.codewars.com/kata/5ae62fcf252e66d44d00008e/go
package main

import (
	"fmt"
	"sort"
)

func ExpressionMatter(a int, b int, c int) int {
	var results = []int{a * b * c, a + b + c, (a + b) * c, a * (b + c)}
	sort.Ints(results)
	return results[3]
}

func main() {
	fmt.Printf("%d", ExpressionMatter(5, 4, 3))
}
