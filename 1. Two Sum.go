package array

// using hashmap to save (target - current_number), it means we just have to
// search for other pair to use this differences value
// O(n) time complexity
// O(n) space complexity
func TwoSum(nums []int, target int) []int {
	// diff map, save the target-curr_num difference
	diff := make(map[int]int, len(nums))

	// for every current number
	// see if the different exist in the map or not
	for i, num := range nums {
		// if the difference is the same as the curr_num
		// we found our pair of result
		if idx, ok := diff[num]; ok {
			return []int{idx, i}
		}
		// if not, let us save the calculation
		diff[target-num] = i
	}
	return nil
}
