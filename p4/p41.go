package main

import (
	"bufio"
	"fmt"
	"os"
)

func count_surrounding_rolls(grid [][]byte, x, y int) int {
	count := 0
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0,  -1},          {0,  1},
		{1,  -1}, {1,  0}, {1,  1},
	}

	for d := range directions {
		surr_x := x + directions[d][0]
		surr_y := y + directions[d][1]
		if is_roll(grid, surr_x, surr_y) {
			count++
		}
	}

	return count
}

func is_roll(grid [][]byte, x, y int) bool {
	if x < 0 || y < 0 || x > len(grid) - 1 || y > len(grid) - 1 {
		return false
	}

	if grid[x][y] == '@' {
		return true
	}
	return false
}

func parse(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func count_accessible_rolls(grid [][]byte) int {
	count := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '.' {
				continue
			} else {
				surr_rolls := count_surrounding_rolls(grid, x, y)
				if surr_rolls < 4 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	grid := parse("input_data.txt")
	result := count_accessible_rolls(grid)
	result_2 := count_accessible_rolls_2(grid)
	fmt.Println("Accessible rolls:", result)
	fmt.Println("Accessible rolls 2:", result_2)
}
