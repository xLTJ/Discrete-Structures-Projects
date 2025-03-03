package lcm

import (
	"testing"
)

func TestFindLCM(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "LCM of 4 and 6",
			a:        4,
			b:        6,
			expected: 12,
		},
		{
			name:     "LCM of 21 and 6",
			a:        21,
			b:        6,
			expected: 42,
		},
		{
			name:     "LCM of 8 and 5",
			a:        8,
			b:        5,
			expected: 40,
		},
		{
			name:     "LCM of coprime numbers 7 and 13",
			a:        7,
			b:        13,
			expected: 91,
		},
		{
			name:     "LCM when one number divides the other",
			a:        6,
			b:        24,
			expected: 24,
		},
		{
			name:     "LCM of identical numbers",
			a:        15,
			b:        15,
			expected: 15,
		},
		{
			name:     "LCM with 1",
			a:        1,
			b:        5,
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindLcm(test.a, test.b)
			if result != test.expected {
				t.Errorf("FindLCM(%d, %d) = %d; expected %d",
					test.a, test.b, result, test.expected)
			}
		})
	}
}
