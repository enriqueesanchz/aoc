// sliding window to check repeated and in order
package main

import (
	"strconv"
)

func check_invalid(s string) bool {
	for window := 1; window < len(s); window++ {
		if len(s) % window != 0 {
			continue
		}

		n := len(s) / window
		var constructed string
		for range n {
			constructed += s[:window]
		}

		if constructed == s {
			return true
		}
	}
	return false
}


func count_invalid_2(a, b int, count int) int {
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)

		if check_invalid(s) {
			count += i
		}
	}
	return count
}
