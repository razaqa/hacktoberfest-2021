package dp

import (
	"github.com/n0irx/go-leet-code/myutil"
)

func MaxSubarrKadane(nums []int) int {
	// initialize the currentMax value and overallMax value
	currentMax, overallMax := 0, 0
	for _, num := range nums {
		currentMax = myutil.MaxInt(0, num, currentMax+num)
		overallMax = myutil.MaxInt(currentMax, overallMax)
	}
	return overallMax
}
