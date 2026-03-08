package leetcode

// Задача №849: Maximize Distance to Closest Person
// https://leetcode.com/problems/maximize-distance-to-closest-person/

func MaxDistToClosest(seats []int) int {
	lastPerson := -1
	maxDist := 0

	for i := range seats {
		if seats[i] == 1 {
			if lastPerson == -1 {
				maxDist = max(maxDist, i)
			} else {
				maxDist = max(maxDist, (i-lastPerson)/2)
			}
			lastPerson = i
		}
	}

	maxDist = max(maxDist, len(seats)-lastPerson-1)

	return maxDist
}
