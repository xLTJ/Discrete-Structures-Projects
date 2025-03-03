package lcm

import (
	"gcd"
	"math"
)

func FindLcm(a, b int) int {
	gcdValue := gcd.FindGcd(a, b)

	return int(math.Abs(float64(a*b))) / gcdValue
}
