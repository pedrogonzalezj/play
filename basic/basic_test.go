package basic

import (
	"testing"
)

func TestAdder(t *testing.T) {
	type test struct {
		a   int
		b   int
		sum int
	}

	tests := []test{
		{a: 1, b: 2, sum: 3},
		{a: 2, b: 2, sum: 4},
		{a: 3, b: 4, sum: 7},
	}

	for _, tc := range tests {
		got := Add(tc.a, tc.b)
		if tc.sum != got {
			t.Errorf("expected '%d' but got '%d'", tc.sum, got)
		}
	}
}
