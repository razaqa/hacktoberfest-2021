package array

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
// Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?

// the idea here is to use bit manipulation here
// for example:
// -> 0, 1, 2, 3, 4
// -> 0, 1, 2, 3
// time complexity: O(n)
// space complexity: O(1)
func MissingNumber(nums []int) int {
	// lenght is n (the total of numbers)
	length := len(nums)

	// this will be accounted to count the expected value if all the list is complete (no missing value)
	expectedXoredValues := length
	// the values we got with the missing value
	gotXoredValues := 0
	for i, val := range nums {
		gotXoredValues ^= val
		expectedXoredValues ^= i
	}

	// this will be the same as subtract all the complete nums with all the nums (with the missing number)
	missingNumber := expectedXoredValues ^ gotXoredValues
	return missingNumber
}
