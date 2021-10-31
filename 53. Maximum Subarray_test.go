package dp

import "testing"

func TestMaxSubarrKadane(t *testing.T) {
	var testCases = []struct {
		nums []int
		want int
	}{
		// Explanation: [4,-1,2,1] has the largest sum = 6.
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, tc := range testCases {
		got := MaxSubarrKadane(tc.nums)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v, numbers input: %v", tc.want, got, tc.nums)
		}
	}
}
