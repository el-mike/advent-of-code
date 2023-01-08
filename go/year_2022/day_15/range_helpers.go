package day_15

import "sort"

func overlap(source, target []int) bool {
	return source[0] <= target[1] && source[1] >= target[0]
}

func testRanges(ranges [][]int) (bool, int) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	openRange := ranges[0]

	for i := 1; i < len(ranges); i += 1 {
		current := ranges[i]

		if !overlap(current, openRange) {
			return false, openRange[1]
		}

		if current[1] > openRange[1] {
			openRange[1] = current[1]
		}
	}

	return true, -1
}
