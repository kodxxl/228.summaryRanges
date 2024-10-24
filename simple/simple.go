package simple

import "fmt"

func SummaryRanges(nums []int) []string {
	result := []string{}
	startCursor, endCursor := 0, 0
	lastIndex := len(nums) - 1

	for i := range nums {
		endCursor = i
		if i < lastIndex {
			if (nums[i] + 1) == nums[i+1] {
				continue
			}
		}
		result = append(result, snapshot(nums, startCursor, endCursor))
		startCursor = (i + 1)
	}
	return result
}

func snapshot(nums []int, startCursor, endCursor int) string {
	if startCursor == endCursor {
		return fmt.Sprintf("%d", nums[startCursor])
	}
	return fmt.Sprintf("%d->%d", nums[startCursor], nums[endCursor])
}
