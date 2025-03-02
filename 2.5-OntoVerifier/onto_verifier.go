package onto_verifier

// IsOnto checks if a given function from {1, 2,...,n} to itself is onto
func IsOnto(n int, function func(int) int) bool {
	codomain := map[int]bool{}
	for i := 1; i <= n; i++ {
		if !codomain[i] {
			codomain[i] = false
		}

		codomain[function(i)] = true
	}

	for _, value := range codomain {
		if !value {
			return false
		}
	}
	return true
}
