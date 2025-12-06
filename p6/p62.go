package main

import (
	"bufio"
	"strconv"
	"strings"
)

func transpose(m [][]byte) [][]byte {
    if len(m) == 0 {
        return [][]byte{}
    }

    rows := len(m)
    cols := len(m[0])
    t := make([][]byte, cols)

	for i := range cols {
        t[i] = make([]byte, rows)
        for j := range rows {
            t[i][j] = m[j][i]
        }
    }
    return t
}

func parse_2(scanner *bufio.Scanner) (nums [][]int, opers []string) {
	m := make([][]byte, 0)
	for scanner.Scan() {
		s := scanner.Text()
		b := []byte(s)
		m = append(m, b)
	}

	o := m[len(m) - 1]
	opers = strings.Fields(string(o))

	m = transpose(m)

	nums = make([][]int, 0)
	row := make([]int, 0)

	for _, r := range m {
		s := string(r)
		s = strings.ReplaceAll(s, " ", "")
		s = strings.Trim(s, "\n")
		if s != "" {
			n, err := strconv.Atoi(s)
			if err == nil {
				row = append(row, n)
			} else {
				num, err := strconv.Atoi(s[:len(s)-1])
				if err != nil {
					panic(err)
				}
				row = append(row, num)
			}
		} else {
			nums = append(nums, row)
			row = make([]int, 0)
		}
	}
	nums = append(nums, row)

	return
}

func calc_2(nums [][]int, opers []string) int {
	if len(nums) != len(opers) {
		panic("len nums != len opers")
	}

	var total int

	for i := range nums {

		var count int
		if opers[i] == "*" {
			count = 1
			for j := range nums[i] {
				count *= nums[i][j]
			}
		} else {
			count = 0
			for j := range nums[i] {
				count += nums[i][j]
			}
		}
		total += count
	}

	return total

}
