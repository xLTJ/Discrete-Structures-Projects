package prime_factorization

import (
	"reflect"
	"testing"
)

func TestPrimeFactorization(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected map[int]int
	}{
		{
			name:     "Number 1",
			input:    1,
			expected: map[int]int{},
		},
		{
			name:     "Prime number",
			input:    13,
			expected: map[int]int{13: 1},
		},
		{
			name:     "Power of 2",
			input:    16,
			expected: map[int]int{2: 4},
		},
		{
			name:     "Product of two primes",
			input:    15,
			expected: map[int]int{3: 1, 5: 1},
		},
		{
			name:     "Multiple prime factors",
			input:    60,
			expected: map[int]int{2: 2, 3: 1, 5: 1},
		},
		{
			name:     "Large number",
			input:    840,
			expected: map[int]int{2: 3, 3: 1, 5: 1, 7: 1},
		},
		{
			name:     "Large prime",
			input:    997,
			expected: map[int]int{997: 1},
		},
		{
			name:     "Product of same prime",
			input:    243,
			expected: map[int]int{3: 5},
		},
		{
			name:     "Large composite",
			input:    123456,
			expected: map[int]int{2: 6, 3: 1, 643: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PrimeFactorization(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("PrimeFactorization(%d) = %v, expected %v",
					test.input, result, test.expected)
			}
		})
	}
}
