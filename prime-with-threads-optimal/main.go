package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var N int32 = 10000000
var CONCURRENCY = 5
// var count int32 = 0

func isPrime(n int32) bool {
	for i := int32(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimesConcurrent(wg *sync.WaitGroup, token, count *int32) {
	defer wg.Done()
	for {
		nextToken := atomic.AddInt32(token, 1)
		if nextToken > N {
			break
		}
		if isPrime(nextToken) {
			atomic.AddInt32(count, 1)
		}
	}
}

func countPrimes() int32 {

	token := int32(-1)
	count := int32(0)
	var wg sync.WaitGroup

	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go countPrimesConcurrent(&wg, &token, &count)
	}

	wg.Wait()
	return count
}

/*

n threads -> run all n threads concurrently
each thread tries to find next available number by incrementing a shared counter
it checks if that number is under the limit and if it is prime
then it increments the shared answer counter

*/

func main() {
	start := time.Now()

	count := countPrimes()

	fmt.Println("Number of primes less than", N, "is", count, "and it took", time.Since(start))
}
