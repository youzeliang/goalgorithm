package _231_power_of_two

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}
