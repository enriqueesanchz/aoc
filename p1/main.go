package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
        panic(err)
    }
    defer file.Close()

	dial1 := 50
	dial2 := 50
	var count1, count2 int

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		direction, n := parse(line)

		dial1 = rotate(dial1, direction, n)
		dial2, count2 = rotate_counting(count2, dial2, direction, n)

		if dial1 == 0 {
			count1 += 1
		}
    }

	fmt.Printf("count1: %d\n", count1)
	fmt.Printf("count2: %d\n", count2)
}
