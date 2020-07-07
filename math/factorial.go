package math

import "math/big"

// FactorialTree Calculate factorial by tree calculation method
func FactorialTree(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}

	if n == 1 {
		return big.NewInt(1)
	}

	if n == 1 || n == 2 {
		return big.NewInt(int64(n))
	}

	return prodTree(2, n)
}

func prodTree(left, right int) *big.Int {
	if left > right {
		return big.NewInt(1)
	}

	if left == right {
		return big.NewInt(int64(left))
	}

	if right-left == 1 {
		return big.NewInt(0).Mul(big.NewInt(int64(left)), big.NewInt(int64(right)))
	}

	m := (left + right) / 2

	return big.NewInt(0).Mul(prodTree(left, m), prodTree(m+1, right))
}
