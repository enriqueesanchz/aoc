package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(line string) []int {
	nums := make([]int, len(line))

	for i, ch := range line {
		nums[i] = int(ch - '0')
	}

	return nums
}

func getMax(nums []int) (index int, n int) {
	n = -1
	for i, num := range nums {
		if num > n {
			n = num
			index = i
		}
	}
	return index, n
}

func getJolts(nums []int) int {
	jolts := -1

	i1, n1 := getMax(nums)
	if i1 == len(nums) - 1 {
		_, n2 := getMax(nums[:i1])
		jolts = n2 * 10 + n1
	} else {
		_, n2 := getMax(nums[i1+1:])
		jolts = n1 * 10 + n2
	}

	if jolts < 0 {
		panic("jolts <0")
	}
	return jolts
}

func main() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	count2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count += getJolts(parse(line))
		jolts2 := maxNumberStack(line, 12)
		count2 += jolts2
	}

	fmt.Println(count)
	fmt.Println(count2)
}
