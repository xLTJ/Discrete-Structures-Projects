package one_to_one

import (
	"testing"
)

// TestIsOneToOne tests the IsOneToOne function with various function implementations
func TestIsOneToOne(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		f        func(int) int
		expected bool
	}{
		{
			name:     "Identity function is one-to-one",
			n:        5,
			f:        func(x int) int { return x },
			expected: true,
		},
		{
			name:     "Square function is one-to-one on positive integers",
			n:        5,
			f:        func(x int) int { return x * x },
			expected: true,
		},
		{
			name:     "Constant function is not one-to-one",
			n:        5,
			f:        func(x int) int { return 42 },
			expected: false,
		},
		{
			name:     "Modulo function is not one-to-one",
			n:        10,
			f:        func(x int) int { return x % 5 },
			expected: false,
		},
		{
			name:     "Custom mapping with all unique values",
			n:        4,
			f:        func(x int) int { return []int{0, 5, 2, 8, 1}[x] },
			expected: true,
		},
		{
			name:     "Custom mapping with duplicate values",
			n:        4,
			f:        func(x int) int { return []int{0, 5, 2, 5, 1}[x] },
			expected: false,
		},
		{
			name:     "Negative output values can be one-to-one",
			n:        4,
			f:        func(x int) int { return []int{0, -3, -2, -1, 0}[x] },
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsOneToOne(tc.n, tc.f)
			if result != tc.expected {
				t.Errorf("IsOneToOne(%v, f) = %v, want %v", tc.n, result, tc.expected)
			}
		})
	}
}
