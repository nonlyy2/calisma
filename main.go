package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	testCases := []struct {
		firstList  [][]int
		secondList [][]int
		expected   [][]int
	}{
		{
			// Стандартный случай: частичные пересечения, точки (5,5), (24,24)
			firstList:  [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			secondList: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			expected:   [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			// Один список пуст
			firstList:  [][]int{{1, 3}, {5, 9}},
			secondList: [][]int{},
			expected:   [][]int{},
		},
		{
			// Один интервал полностью поглощает другой
			firstList:  [][]int{{1, 10}},
			secondList: [][]int{{4, 5}, {7, 8}},
			expected:   [][]int{{4, 5}, {7, 8}},
		},
		{
			// Интервалы касаются только границами
			firstList:  [][]int{{1, 2}, {3, 4}},
			secondList: [][]int{{2, 3}},
			expected:   [][]int{{2, 2}, {3, 3}},
		},
	}

	fmt.Println("📅 Тестируем Interval List Intersections (#986)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.IntervalIntersection(tc.firstList, tc.secondList)

		// Для вывода в консоль при ошибке
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден\n", i+1)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка!\n   First:  %v\n   Second: %v\n   Ожидали: %v\n   Получили: %v\n",
				i+1, tc.firstList, tc.secondList, tc.expected, result)
		}
	}
}
