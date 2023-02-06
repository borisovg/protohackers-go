package prime

import (
	"math/big"
	"testing"
)

func TestIsPrime(t *testing.T) {
	type test struct {
		num  *big.Int
		want bool
	}

	tests := []test{
		{num: big.NewInt(1), want: false},
		{num: big.NewInt(2), want: true},
		{num: big.NewInt(8), want: false},
		{num: big.NewInt(17), want: true},
	}

	for _, tc := range tests {
		got := IsPrime(tc.num)
		if tc.want != got {
			t.Fatalf("for value: %v, expected: %v, got: %v", tc.num, tc.want, got)
		}
	}
}
