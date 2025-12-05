package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse_ranges(scanner *bufio.Scanner) [][2]int64 {
	ranges_array := make([][2]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		start, end := parse_range_to_int(line)
		ranges_array = append(ranges_array, [2]int64{start, end})
	}
	return ranges_array
}

func count_fresh(scanner *bufio.Scanner, ranges_array [][2]int64) int {
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		var num int64
		fmt.Sscanf(line, "%d", &num)

		for _, r := range ranges_array {
			if num >= r[0] && num <= r[1] {
				count++
				break
			}
		}

	}
	return count
}

func parse_range_to_int(line string) (int64, int64) {
	var start, end int64
	fmt.Sscanf(line, "%d-%d", &start, &end)
	return start, end
}

func main() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges_array := parse_ranges(scanner)
	count := count_fresh(scanner, ranges_array)
	fmt.Println(count)

	sorted_ranges := sort_ranges(ranges_array)
	merged_ranges := merge_ranges(sorted_ranges)
	fmt.Println(count_fresh_total(merged_ranges))
}
