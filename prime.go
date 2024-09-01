package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeGenerator(n int, ch chan int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go primeGenerator(30, ch)

	for prime := range ch {
		fmt.Println(prime)
	}
}
