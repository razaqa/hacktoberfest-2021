package array

// using hashset to save iterated num
// time complexity: O(n)
// space complexity: O(n)
func ContainsDuplicate(nums []int) bool {
	set := map[int]struct{}{}

	for _, num := range nums {
		// if current num is already in set, there is a duplication
		if _, ok := set[num]; ok {
			return true
		}

		// if not we set the num true in the set
		set[num] = struct{}{}
	}

	// if after iterating all the num, there is no duplication
	// there is no duplication
	return false
}
