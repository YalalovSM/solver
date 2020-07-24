package math

import (
	"log"
	"math/big"
	"sync"
	"time"
)

// FactorialTree Calculate factorial by tree calculation method
func FactorialTree(n int) *big.Int {
	defer timeTrack(time.Now(), n)
	// number of goroutines to run in parallel
	workers := 2

	if n < 0 {
		return big.NewInt(0)
	}

	if n == 1 || n == 2 {
		return big.NewInt(int64(n))
	}

	if n < workers+1 {
		return prodTree(2, n)
	}

	ch := make(chan *big.Int, workers)

	wg := &sync.WaitGroup{}

	diff := (n - 2) / workers
	left := 2
	var right int

	for i := 0; i < workers; i++ {
		if right = left + diff; right > n {
			right = n
		}

		wg.Add(1)
		go goProdTree(left, right, ch, wg)

		if left = right + 1; left > n {
			break
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	f := big.NewInt(1)

	for v := range ch {
		f.Mul(f, v)
	}

	return f
}

func goProdTree(left, right int, ch chan *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()

	res := prodTree(left, right)
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

func timeTrack(start time.Time, n int) {
	elapsed := time.Since(start)
	log.Printf("calculation of %d! took %s", n, elapsed)
}
