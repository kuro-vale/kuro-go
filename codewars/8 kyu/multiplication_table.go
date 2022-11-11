// https://www.codewars.com/kata/5a2fd38b55519ed98f0000ce/go
package main

import (
	"fmt"
)

func MultiTable(number int) string {
	var table = ""
	for i := 1; i <= 10; i++ {
		table += fmt.Sprintf("%d * %d = %d", i, number, i*number)
		if i < 10 {
			table += "\n"
		}
	}
	return table
}

func main() {
	fmt.Printf("%s", MultiTable(5))
}
