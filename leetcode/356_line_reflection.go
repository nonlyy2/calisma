package leetcode

import (
	"math"
)

// Задача №356: Line Reflection
// https://leetcode.com/problems/line-reflection/

func IsReflected(points [][]int) bool {
	minX, maxX := math.MaxInt32, math.MinInt32

	pointSet := make(map[[2]int]struct{})

	for _, p := range points {
		x, y := p[0], p[1]

		minX = min(minX, x)
		maxX = max(maxX, x)

		pointSet[[2]int{x, y}] = struct{}{}
	}

	targetSum := minX + maxX

	for p := range pointSet {
		x, y := p[0], p[1]

		mirrorPoint := [2]int{targetSum - x, y}

		if _, found := pointSet[mirrorPoint]; !found {
			return false
		}
	}

	return true
}
