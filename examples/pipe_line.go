package examples

import (
	"fmt"
	"math"
)

// isPrime returns true if n is a prime number.
func isPrime(n int) bool {
	// Handle numbers less than 2 (0 and 1 are not prime)
	if n < 2 {
		return false
	}

	// Check for divisors from 2 up to the square root of n.
	// If any divisor is found, n is not prime.
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFilter(muns []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, mun := range muns {
			if isPrime(mun) {
				out <- mun
			}
		}
		close(out)
	}()

	return out
}

func Pow(muns <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for m := range muns {
			out <- m * m
		}
		close(out)
	}()

	return out
}

func RunPipeLineExample() {
	nums := []int{1, 2, 3, 4, 5}

	primeNums := primeFilter(nums)

	rs := Pow(primeNums)

	for num := range rs {
		fmt.Println(num)
	}
}
