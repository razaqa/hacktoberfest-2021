package dp

import "testing"

func TestMissingNumber(t *testing.T) {
	var testCases = []struct {
		input int
		want  int
	}{
		// Explanation: There are two ways to climb to the top.
		// 1. 1 step + 1 step
		// 2. 2 steps
		{2, 2},

		// Explanation: There are three ways to climb to the top.
		// 1. 1 step + 1 step + 1 step
		// 2. 1 step + 2 steps
		// 3. 2 steps + 1 step
		{3, 3},
	}

	for _, tc := range testCases {
		got := make([]int, 0)

		got = append(got, ClimbStairs(tc.input))
		got = append(got, ClimbStairsMemoizationTopDown(tc.input))
		got = append(got, ClimbStairsNaiveRecursive(tc.input))
		got = append(got, ClimbStairsTabulationBottomUp(tc.input))

		for i := 1; i < len(got); i++ {
			if got[i] != got[0] {
				t.Fatal("The return value in each method should be the same")
			}
		}

		if got[0] != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
