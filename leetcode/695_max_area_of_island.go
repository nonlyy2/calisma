package leetcode

// Задача №695: Max Area of Island
// https://leetcode.com/problems/max-area-of-island/

func MaxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				maxArea = max(maxArea, calculateArea(grid, r, c))
			}
		}
	}

	return maxArea
}

func calculateArea(grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0

	return 1 + calculateArea(grid, r+1, c) + calculateArea(grid, r-1, c) + calculateArea(grid, r, c+1) + calculateArea(grid, r, c-1)
}
