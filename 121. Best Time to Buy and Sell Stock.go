package dp

// time complexity = O(n)
// space complexity = O(1)
// we track a window of difference of the prices
// when we add the prices and it adds to less than zero, we reset the calculation
func MaxProfit(prices []int) int {
	// tempMax will hold current window max
	// resMax will hold overall max
	tempMax, resMax := 0, 0

	for i := 1; i < len(prices); i++ {
		tempMax += prices[i] - prices[i-1]
		// we just want the positive value
		if tempMax < 0 {
			tempMax = 0
		}
		// if tempMax bigger than resMax, update value
		if tempMax > resMax {
			resMax = tempMax
		}
	}
	return resMax
}
