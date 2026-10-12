package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var N = 10000000
var CONCURRENCY = 5
// var count int32 = 0

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimesBatch(start, end int, wg *sync.WaitGroup, count *int32) {
	defer wg.Done()
	for i := start; i < end; i++ {
		if isPrime(i) {
			atomic.AddInt32(count, 1)  // This is thread-safe and avoids race conditions
			// count++  // This is not thread-safe and can lead to race conditions
		}
	}
}

func countPrimes(n int) int32 {

	count := int32(0)
	var wg sync.WaitGroup

	batchSize := n / CONCURRENCY
	for i := 0; i < CONCURRENCY; i++ {
		start := i * batchSize
		end := start + batchSize
		if i == CONCURRENCY-1 {
			end = n
		}

		wg.Add(1)
		go countPrimesBatch(start, end, &wg, &count)
	}

	wg.Wait()
	return count
}

func main() {
	start := time.Now()

	count := countPrimes(N)

	fmt.Println("Number of primes less than", N, "is", count, "and it took", time.Since(start))
}
