package prime_checker

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "1 is not prime",
			input:    1,
			expected: false,
		},
		{
			name:     "2 is prime",
			input:    2,
			expected: true,
		},
		{
			name:     "3 is prime",
			input:    3,
			expected: true,
		},
		{
			name:     "4 is not prime",
			input:    4,
			expected: false,
		},
		{
			name:     "5 is prime",
			input:    5,
			expected: true,
		},
		{
			name:     "6 is not prime",
			input:    6,
			expected: false,
		},
		{
			name:     "7 is prime",
			input:    7,
			expected: true,
		},
		{
			name:     "8 is not prime",
			input:    8,
			expected: false,
		},
		{
			name:     "9 is not prime",
			input:    9,
			expected: false,
		},
		{
			name:     "10 is not prime",
			input:    10,
			expected: false,
		},
		{
			name:     "11 is prime",
			input:    11,
			expected: true,
		},
		{
			name:     "13 is prime",
			input:    13,
			expected: true,
		},
		{
			name:     "15 is not prime",
			input:    15,
			expected: false,
		},
		{
			name:     "17 is prime",
			input:    17,
			expected: true,
		},
		{
			name:     "19 is prime",
			input:    19,
			expected: true,
		},
		{
			name:     "23 is prime",
			input:    23,
			expected: true,
		},
		{
			name:     "25 is not prime",
			input:    25,
			expected: false,
		},
		{
			name:     "29 is prime",
			input:    29,
			expected: true,
		},
		{
			name:     "97 is prime",
			input:    97,
			expected: true,
		},
		{
			name:     "99 is not prime",
			input:    99,
			expected: false,
		},
		{
			name:     "101 is prime",
			input:    101,
			expected: true,
		},
		{
			name:     "Large prime: 7919",
			input:    7919,
			expected: true,
		},
		{
			name:     "Large non-prime: 7917",
			input:    7917,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPrime(tc.input)
			if got != tc.expected {
				t.Errorf("IsPrime(%d) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}

// Benchmark to test performance
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(7919) // Test with a large prime
	}
}
