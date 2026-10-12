package main

import (
	"fmt"
	"time"
)

var N = 10000000

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	start := time.Now()

	count := countPrimes(N)

	fmt.Println("Number of primes less than", N, "is", count, "and it took", time.Since(start))
}
