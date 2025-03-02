package one_to_one

// IsOneToOne checks if a function is one-to-one
func IsOneToOne(n int, function func(int) int) bool {
	var codomain = map[int]bool{}
	for i := 1; i <= n; i++ {
		_, exists := codomain[function(i)]

		if exists {
			return false
		}

		codomain[function(i)] = true
	}
	return true
}
