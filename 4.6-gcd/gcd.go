package gcd

func FindGcd(a, b int) int {
	if a == b {
		return a
	}

	// make sure a >= b
	x := max(a, b)
	y := min(a, b)

	for y != 0 {
		rest := x % y
		x = y
		y = rest
	}

	return x
}
