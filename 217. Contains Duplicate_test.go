package array

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {

	var testCases = []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range testCases {
		got := ContainsDuplicate(tc.nums)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
