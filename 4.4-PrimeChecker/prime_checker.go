package prime_checker

import "math"

func IsPrime(n int) bool {
	// 2 is an edge case and is a prime number
	if n == 2 {
		return true
	}

	// integers that are less than 2, or divisible by 2 are not prime numbers
	if n < 2 || n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
