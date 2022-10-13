func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	if n%2 != 0 {
		return x * myPow(x*x, n/2)
	}

	return myPow(x*x, n/2)
}