// https://www.codewars.com/kata/5bb904724c47249b10000131/go

package main

import "strings"

func main() {
	print(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))
}

func Points(games []string) int {
	var score = 0
	for i := 0; i < len(games); i++ {
		var points = strings.Split(games[i], ":")
		if points[0] > points[1] {
			score += 3
		} else if points[0] == points[1] {
			score += 1
		}
	}
	return score
}
