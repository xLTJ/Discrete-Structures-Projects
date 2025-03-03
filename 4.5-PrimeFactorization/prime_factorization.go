package prime_factorization

// first 10 prime numbers, probably wont need more but just add more if needed
var primeNumbers = [15]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

func PrimeFactorization(n int) map[int]int {
	factorization := map[int]int{}

	// empty if 1 or less
	if n <= 1 {
		return factorization
	}

	// check divisibility by 2
	for n%2 == 0 {
		factorization[2]++
		n /= 2
	}

	// check divisibility by the pre-computed primes
	for _, prime := range primeNumbers {
		for n%prime == 0 {
			factorization[prime]++
			n /= prime
		}
	}

	// for large numbers, check divisibility by further numbers
	for i := 47; n != 1; i += 2 {
		for n%i == 0 {
			factorization[i]++
			n /= i
		}
	}

	return factorization
}
