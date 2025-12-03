package main

import "strconv"

func maxNumberStack(s string, k int) int {
	stack := []byte{}
	toRemove := len(s) - k

	for i := 0; i < len(s); i++ {
		digit := s[i]
		// Pop smaller digits from stack if we can still remove
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1] // pop
			toRemove--
		}
		stack = append(stack, digit)
	}

	// Truncate stack if it's longer than k
	if len(stack) > k {
		stack = stack[:k]
	}

	sol, err := strconv.Atoi(string(stack))
	if err != nil {
		panic(err)
	}
	return sol
}

