package dp

import "testing"

func TestMaxProfit(t *testing.T) {
	var testCases = []struct {
		nums []int
		want int
	}{
		// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
		// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.s
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range testCases {
		got := MaxProfit(tc.nums)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v, numbers input: %v", tc.want, got, tc.nums)
		}
	}
}
