package dp

// 1. Naive Recursive
// this is the most intuitive solution
// but the time complexity will be O(2^n),
// and the space complexity will be  O(2^n) recrusive call stack
func ClimbStairsNaiveRecursive(n int) int {
	// recrusive helper function for climbstairs
	var rec func(n int) int
	rec = func(n int) int {
		// base cases
		if n == 0 || n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}

		return rec(n-1) + rec(n-2)
	}

	// calling the helper function
	return rec(n)
}

// 2. Top Down Recursive with Memoization
// better version than the Naive Recursive solution
// O(n) time complexity
// O(n) space complexity
func ClimbStairsMemoizationTopDown(n int) int {
	// memoization variable
	memo := make(map[int]int)

	// recrusive helper function for climbstairs
	var rec func(n int) int
	rec = func(n int) int {
		// base cases
		if n == 0 || n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}

		// if the calculation of f(n) already exist in the memo, return it
		if result, ok := memo[n]; ok {
			return result
		}

		// save the calculation of f(n)
		memo[n] = rec(n-1) + rec(n-2)
		return memo[n]
	}

	// calling the helper function
	return rec(n)
}

// 3. Bottom Up with Tabulation
// Iterative version to reduce recursive call stack
// O(n) time complexity
// O(n) space complexity
func ClimbStairsTabulationBottomUp(n int) int {
	// initialize dp tabulation
	dp := make([]int, n+1)

	// base case
	dp[0] = 1
	dp[1] = 1

	// fill the tabulation with the value from two previous calculation
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// return the last index of tabulation
	return dp[n]
}

// 4. Bottom Up with O(1) space complexity
// Better version of 3. with using O(1) space complexity
// O(n) time complexity
// O(1) space complexity
func ClimbStairs(n int) int {
	// initialize dp tabulation
	prev, next := 1, 1

	// fill the tabulation with the value from two previous calculation
	for i := 2; i < n+1; i++ {
		prev, next = next, prev+next
	}

	return next
}
