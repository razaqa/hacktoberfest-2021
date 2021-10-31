package bitman

// we will use bitwise and operation to determine if the rightmost value is 1
// if so, we raise the counter, after that we shift right the bits of num
func hammingWeight(num uint32) int {
	// count will hold the number of 1 in the bits
	count := 0
	for num != 0 {
		// we add count if num&1 is 1, (the last bit is 1)
		// after that we shift right num
		count, num = count+int((num&1)), num>>1
	}
	return count
}
