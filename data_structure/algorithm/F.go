package twostack2queue

// F 斐波那契
func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}

	return F(n-1, a2, a1+a2)

}
