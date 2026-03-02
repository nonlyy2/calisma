package additional

// Задача: Слияние интервалов (Merge Intervals)
// Часто встречается на собеседованиях в Яндекс (Бэкенд).
//
// Дан массив интервалов, где intervals[i] = [start, end].
// Нужно объединить все перекрывающиеся интервалы и вернуть массив
// неперекрывающихся интервалов, которые охватывают все входные интервалы.
//
// Ограничения:
// - 1 <= intervals.length <= 10^4
// - intervals[i].length == 2
// - 0 <= start <= end <= 10^4
// - Входной массив может быть НЕ отсортирован.
//
// Пример 1:
// Вход: [[1,3],[2,6],[8,10],[15,18]]
// Выход: [[1,6],[8,10],[15,18]]
//
// Пример 2:
// Вход: [[1,4],[4,5]]
// Выход: [[1,5]]
//
// Пример 3:
// Вход: [[1,10], [2,3], [4,5]]
// Выход: [[1,10]]

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	current := intervals[0]

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] <= current[1] {
			current[1] = max(current[1], intervals[i][1])
		} else {
			res = append(res, current)
			current = intervals[i]
		}
	}

	res = append(res, current)

	return res
}
