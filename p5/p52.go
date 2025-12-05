package main

import (
	"sort"
)

func sort_ranges(ranges [][2]int64) [][2]int64 {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	return ranges
}

func merge_ranges(ranges [][2]int64) [][2]int64 {
	var merged [][2]int64

	for _, r := range ranges {
		lo, hi := r[0], r[1]

		if len(merged) == 0 {
			merged = append(merged, [2]int64{lo, hi})
			continue
		}

		last := merged[len(merged)-1]
		if lo <= last[1]+1 {
			if hi > last[1] {
				merged[len(merged)-1][1] = hi
			}
		} else {
			merged = append(merged, [2]int64{lo, hi})
		}
	}
	
	return merged
}

func count_fresh_total(merged_ranges [][2]int64) int64 {
	var count int64

	for _, r := range merged_ranges {
		if r[0] != -1 && r[1] != -1 {
			count += r[1] - r[0] + 1
		}
	}
	return count
}
