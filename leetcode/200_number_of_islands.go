package leetcode

// Задача №200: Number of Islands
// https://leetcode.com/problems/number-of-islands/

func NumIslands(grid [][]byte) int {
	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				count++
				recursiveHelper(grid, r, c)
			}
		}
	}

	return count
}

func recursiveHelper(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'

	recursiveHelper(grid, r+1, c)
	recursiveHelper(grid, r-1, c)
	recursiveHelper(grid, r, c+1)
	recursiveHelper(grid, r, c-1)
}
