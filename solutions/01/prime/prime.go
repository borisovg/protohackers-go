package prime

import "math/big"

func IsPrime(num *big.Int) bool {
	return num.ProbablyPrime(41)
}
