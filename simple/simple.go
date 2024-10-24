package simple

import "strconv"

func SummaryRanges(nums []int) []string {
	result := []string{}
	startCursor, endCursor := 0, 0
	lenNums := len(nums)

	for i := range nums {
		endCursor = i
		if i < (lenNums - 1) {
			if (nums[i] + 1) == nums[i+1] {
				continue
			}
		}
		annotation := makeAnnotation(nums, startCursor, endCursor)
		result = append(result, annotation)
		startCursor = (i + 1)
	}

	return result
}

func makeAnnotation(nums []int, startCursor, endCursor int) string {
	first := nums[startCursor]
	last := nums[endCursor]

	if startCursor == endCursor {
		return strconv.Itoa(first)
	}
	return strconv.Itoa(first) + "->" + strconv.Itoa(last)
}
