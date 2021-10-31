package dp

import "math"

// Top Down approach with memoization
func CoinChangeTopDownMemoization(coins []int, amount int) int {
	// initialize recrusive helper
	var coinChangeHelper func(amount int) int
	// initizlie memo variable
	memo := make([]int, amount+1)
	// method helper for recursive
	coinChangeHelper = func(amount int) int {
		// base cases
		if amount == 0 || memo[amount] != 0 {
			return memo[amount]
		}
		// set the highest value for default value
		res := math.MaxInt64
		for _, coin := range coins {
			// if the amount is the same as coin, just add one and return
			if coin == amount {
				res = 1
			} else if coin < amount {
				// if not, we have to configure all the different answer
				// for each coin available, and we find the minimum value
				// recursive call all the coin that can be used for current amount
				tmp := 1 + coinChangeHelper(amount-coin)
				if tmp != 0 && tmp < res {
					res = tmp
				}
			}
		}
		// case if no coin can perform to make number of needed amount
		if res == math.MaxInt64 {
			memo[amount] = -1
			return -1
		}
		// save the calculation to memo
		memo[amount] = res
		return memo[amount]
	}
	// call helper function here
	return coinChangeHelper(amount)
}
