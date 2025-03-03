package gcd

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Small numbers",
			a:        12,
			b:        8,
			expected: 4,
		},
		{
			name:     "Larger numbers",
			a:        48,
			b:        18,
			expected: 6,
		},
		{
			name:     "Relatively prime numbers",
			a:        101,
			b:        103,
			expected: 1,
		},
		{
			name:     "One number divides the other",
			a:        25,
			b:        5,
			expected: 5,
		},
		{
			name:     "Equal numbers",
			a:        17,
			b:        17,
			expected: 17,
		},
		{
			name:     "Classic example",
			a:        1071,
			b:        462,
			expected: 21,
		},
		{
			name:     "Reversed order of inputs",
			a:        24,
			b:        60,
			expected: 12,
		},
		{
			name:     "Large numbers",
			a:        3192,
			b:        2760,
			expected: 24,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindGcd(test.a, test.b)
			if result != test.expected {
				t.Errorf("FindGcd(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

// Test that the function handles arguments in either order
func TestGCDCommutative(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{48, 18},
		{101, 103},
		{1071, 462},
		{3192, 2760},
	}

	for _, test := range tests {
		t.Run("Commutative property", func(t *testing.T) {
			result1 := FindGcd(test.a, test.b)
			result2 := FindGcd(test.b, test.a)
			if result1 != result2 {
				t.Errorf("FindGcd(%d, %d) = %d, but FindGcd(%d, %d) = %d",
					test.a, test.b, result1, test.b, test.a, result2)
			}
		})
	}
}
