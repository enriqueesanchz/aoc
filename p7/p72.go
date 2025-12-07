package main

func memo_dfs(grid [][]byte, x, y int, count int, memo map[[2]int]int) int {
	key := [2]int{x, y}
	if val, ok := memo[key]; ok {
		return val
	}

	for grid[x][y] == '.' {
		x += 1
		if x == len(grid) {
			return 1
		}
	}
	res := memo_dfs(grid, x, y-1, count, memo) + memo_dfs(grid, x, y+1, count, memo)
	memo[key] = res
	return res
}
