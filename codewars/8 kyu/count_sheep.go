// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/go

package main

import "fmt"

func countSheep(num int) string {
	var sheep = ""
	for i := 1; i <= num; i++ {
		sheep += fmt.Sprintf("%d sheep...", i)
	}
	return sheep
}

func main() {
	print(countSheep(2))
}
