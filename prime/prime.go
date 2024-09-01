package prime

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

func PrimeGenerator(n int, CH chan int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			CH <- i
		}
	}
	close(CH)
}
