package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
        panic(err)
    }
    defer file.Close()

	dial := 50
	var count int

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		direction, n := parse(line)
		dial = rotate(dial, direction, n)

		if dial == 0 {
			count += 1
		}
    }

	fmt.Printf("count: %d\n", count)
}
