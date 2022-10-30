// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/go

package kata

import "math"

func QuarterOf(month int) int {
  return int(math.Ceil(float64(month)/3))
}
