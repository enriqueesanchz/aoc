package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func parse(scanner *bufio.Scanner) ([][]int, []string) {
	nums := make([][]int, 0)
	opers := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		fields := strings.Fields(s)

		n := make([]int, 0)
		for _, f := range fields {
			v, err := strconv.Atoi(f)
			if err == nil {
				n = append(n, v)
			}

		}
		if len(n) != 0 {
			nums = append(nums, n)
		} else {
			opers = fields
		}
	}

	return nums, opers
}

func calc(nums [][]int, opers []string) int {
	if len(nums[0]) != len(opers) {
		panic("len nums != len opers")
	}

	var total int

	for i := range nums[0] {

		var count int
		if opers[i] == "*" {
			count = 1
			for j := range nums {
				count *= nums[j][i]
			}
		} else {
			count = 0
			for j := range nums {
				count += nums[j][i]
			}
		}
		total += count
	}

	return total
}

func p1() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	start := time.Now()
	nums, opers := parse(scanner)
	fmt.Printf("[%v] %d\n", time.Since(start), calc(nums, opers))
}

func p2() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	start := time.Now()
	nums, opers := parse_2(scanner)
	fmt.Printf("[%v] %d\n", time.Since(start), calc_2(nums, opers))
}

func main() {
	fmt.Println("[+] day 6")
	p1()
	p2()
}
