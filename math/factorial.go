package math

import (
	"fmt"
	"log"
	"math/big"
)

// FactorialTree Calculate factorial by tree calculation method
func FactorialTree(n int) *big.Int {
	threads := 2

	if n < 0 {
		return big.NewInt(0)
	}

	if n == 1 {
		return big.NewInt(1)
	}

	if n == 1 || n == 2 {
		return big.NewInt(int64(n))
	}

	if n < threads+1 {
		return prodTree(2, n)
	}

	// TODO replace the code below which works only for even N
	ch := make(chan *big.Int)
	upper := n / threads
	log.Printf("Before first goroutine left = %d, right = %d", 2, upper)
	go goProdTree(2, upper, ch)

	for i := 1; i < threads; i++ {
		left := (n/threads)*i + 1
		right := (n / threads) * (i + 1)
		log.Printf("Before goroutine left = %d, right = %d, i = %d", left, right, i)
		go goProdTree(left, right, ch)
	}

	// TODO the code below had to return the calculated factorial
	// need to read the golang doc about processing the messages when multiple senders produce
	// the messages into one channel
	f := big.NewInt(1)
	for {
		select {
		case v := <-ch:
			fmt.Printf("v = %s\n", v.String())
			f.Mul(f, v)
		}
	}

	// return f
}

func goProdTree(left, right int, ch chan *big.Int) {
	res := prodTree(left, right)
	log.Printf("res = %s\n", res.String())
	ch <- res
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
