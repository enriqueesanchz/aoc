package main

import (
	"bufio"
	"fmt"
	"os"
)

func p1() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := read_grid(scanner)
	count := beam(grid, 0, 70, 0)
	fmt.Println(count)
}

func p2() {
	file, err := os.Open("input_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := read_grid(scanner)
	memo := make(map[[2]int]int)
	count_2 := memo_dfs(grid, 0, 70, 0, memo)
	fmt.Println(count_2)
}

func main() {
	p1()
	p2()
}

func read_grid(scanner *bufio.Scanner) [][]byte {
	grid := make([][]byte, 0)

	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())

		grid = append(grid, line)
	}

	return grid
}

func beam(grid [][]byte, x, y int, count int) int {
	if y < 0 || y >= len(grid[0]) {
		return count
	}
	if x >= len(grid) {
		return count
	}

	if grid[x][y] == '|' {
		return count
	}

	if grid[x][y] == '^' {
		count++
		count = beam(grid, x, y-1, count)
		count = beam(grid, x, y+1, count)
		return count
	} else {
		grid[x][y] = '|'
		count = beam(grid, x+1, y, count)
		return count
	}
}
