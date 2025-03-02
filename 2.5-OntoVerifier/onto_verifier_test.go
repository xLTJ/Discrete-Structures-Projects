package onto_verifier

import (
	"testing"
)

// TestIsOnto tests the IsOnto function with various function implementations
func TestIsOnto(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		f        func(int) int
		expected bool
	}{
		{
			name:     "Identity function is onto",
			n:        5,
			f:        func(x int) int { return x },
			expected: true,
		},
		{
			name:     "Constant function is not onto for n > 1",
			n:        5,
			f:        func(x int) int { return 3 },
			expected: false,
		},
		{
			name:     "Constant function is onto for n = 1",
			n:        1,
			f:        func(x int) int { return 1 },
			expected: true,
		},
		{
			name:     "Permutation function is onto",
			n:        5,
			f:        func(x int) int { return []int{0, 5, 4, 3, 2, 1}[x] },
			expected: true,
		},
		{
			name:     "Function missing output value 3",
			n:        5,
			f:        func(x int) int { return []int{0, 1, 2, 4, 5, 5}[x] },
			expected: false,
		},
		{
			name:     "Function with all outputs = inputs % n + 1",
			n:        5,
			f:        func(x int) int { return x%5 + 1 },
			expected: true,
		},
		{
			name:     "Function mapping to only even numbers",
			n:        6,
			f:        func(x int) int { return x * 2 % 6 },
			expected: false,
		},
		{
			name:     "Larger domain with missing values",
			n:        10,
			f:        func(x int) int { return (x % 5) + 1 },
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsOnto(tc.n, tc.f)
			if result != tc.expected {
				t.Errorf("IsOnto(%v, f) = %v, want %v", tc.n, result, tc.expected)
			}
		})
	}
}
