package bitman

// when we trace the addition of two number
// it's basically a binary operation xor (^) and bitwise and (&)
// xor is used for getting sum of two number, the bitwise and is used to find carry
// time complexity = O(b)
// space complexity = O(n)
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
