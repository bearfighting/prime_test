package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(num int32) bool {
	if num == 2 {
		return true
	}
	if num == 1 || num%2 == 0 {
		return false
	}
	to := int32(math.Sqrt(float64(num)))
	for div := int32(3); div <= to; div += 2 {
		if num%div == 0 {
			return false
		}
	}
	return true
}

func do(N int32) {
	for i := int32(0); i < N; i++ {
		prime := isPrime(i)
		if prime {
			// fmt.Printf("%+v: %+v\n", i, prime)
		}
	}
}

func main() {
	st := time.Now()
	do(10_000_000)
	fmt.Printf("%+v\n", time.Since(st))
}
