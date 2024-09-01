package main

import (
	"GO/prime"
	"fmt"
)

func sum(numbers []int, ch chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	ch <- total
}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	mid := len(numbers) / 2
	first := numbers[:mid]
	second := numbers[mid:]

	ch := make(chan int)

	go sum(first, ch)
	go sum(second, ch)

	sum1 := <-ch
	sum2 := <-ch
	total := sum1 + sum2

	fmt.Printf("Total: %d\n\n", total)

	fmt.Println("First 10 Prime Numbers :")

	CH := make(chan int)

	go prime.PrimeGenerator(30, CH)

	for primeNum := range CH {
		fmt.Print(primeNum)
		fmt.Print(" ")
	}
}
