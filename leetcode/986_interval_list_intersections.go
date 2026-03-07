package leetcode

// Задача №986: Interval List Intersections
// https://leetcode.com/problems/interval-list-intersections/

func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 && len(secondList) == 0 {
		return [][]int{}
	}

	res := [][]int{}

	i := 0
	j := 0

	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		if end >= start {
			res = append(res, []int{start, end})
		}

		if firstList[i][1] > secondList[j][1] {
			j++
		} else {
			i++
		}
	}

	return res
}
