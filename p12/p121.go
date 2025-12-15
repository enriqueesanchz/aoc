package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("[p1] %d\n", p1())
}

func p1() int {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	shapes := parse_shapes(scanner)

	hasText := true
	count := 0
	for hasText {
		s := scanner.Text()
		
		dims := strings.Split(s, ": ")
		dim_x, err:= strconv.Atoi(strings.Split(dims[0], "x")[0])
		if err != nil {
			panic(err)
		}
		dim_y, err := strconv.Atoi(strings.Split(dims[0], "x")[1])
		if err != nil {
			panic(err)
		}

		max_size := dim_x * dim_y

		size := 0
		for i, count := range strings.Fields(dims[1]) {
			count_int, err := strconv.Atoi(count)
			if err != nil {
				panic(err)
			}

			size += count_int * shapes[i]
		}

		if size <= max_size {
			count++
		}
		
		hasText = scanner.Scan()
	}
	return count
}

func parse_shapes(scanner *bufio.Scanner) map[int]int {
	shapes := make(map[int]int, 0)

	for scanner.Scan() {
		s := scanner.Text()
		if strings.ContainsAny(s, "x") {
			break
		}

		idx, err := strconv.Atoi(s[:1])
		if err != nil {
			panic(err)
		}

		size := 0
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				break
			}
			for _, c := range text {
				if c == '#' {
					size++
				}
			}
		}
		shapes[idx] = size
	}

	return shapes
}
