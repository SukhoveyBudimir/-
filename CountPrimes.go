func countPrimes(n int) int {
	primes := make([]bool, n)
	result := 0
	for i := 2; i < n; i++ {
		primes[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if primes[i] {
			for j := i + i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	for i := 2; i < n; i++ {
		if primes[i] {
			result++
		}
	}
	return result
}