package piscine

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

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}
