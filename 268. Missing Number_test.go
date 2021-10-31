package array

import "testing"

func TestMissingNumber(t *testing.T) {
	var testCases = []struct {
		nums []int
		want int
	}{
		// Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3].
		// 2 is the missing number in the range since it does not appear in nums.
		{[]int{3, 0, 1}, 2},

		// Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2].
		// 2 is the missing number in the range since it does not appear in nums.
		{[]int{0, 1}, 2},

		// Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9].
		// 8 is the missing number in the range since it does not appear in nums.
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tc := range testCases {
		got := MissingNumber(tc.nums)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
