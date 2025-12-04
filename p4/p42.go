package main

func count_accessible_rolls_2(grid [][]byte) int {
	count := 0
	last_count := 0
	for range grid {
		for x := range grid {
			for y := range grid[x] {
				if grid[x][y] == '.' {
					continue
				} else {
					surr_rolls := count_surrounding_rolls(grid, x, y)
					if surr_rolls < 4 {
						count++
						grid[x][y] = '.'
					}
				}
			}
		}
		if count == last_count {
			break
		}
		last_count = count
	}
	return count
}
