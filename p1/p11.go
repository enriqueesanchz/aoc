package main

import (
	"strconv"
)

func parse(line string) (byte, int) {
	n, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}

	return line[0], n
}

func rotate(dial int, direction byte, n int) int {
	if direction == 'R' {
		dial = (dial + n) % 100
	} else {
		dial = (dial - n) % 100
		if dial < 0 {
			dial += 100
		}
	}
	return dial
}

