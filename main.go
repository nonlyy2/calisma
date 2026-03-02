package main

import (
	"fmt"
	"neetcode/additional" // Поменяй на свой путь импорта
	"reflect"
)

func main() {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "Обычный случай",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "Стык в стык",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "Несортированный вход",
			input:    [][]int{{15, 18}, {1, 3}, {8, 10}, {2, 6}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "Один большой интервал поглощает маленькие",
			input:    [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}},
			expected: [][]int{{1, 10}},
		},
		{
			name:     "Пустой массив",
			input:    [][]int{},
			expected: [][]int{},
		},
	}

	fmt.Println("🚀 Начинаем тестирование Merge Intervals...")

	for _, tc := range tests {
		// Копируем вход, так как сортировка может изменить исходный слайс
		inputCopy := make([][]int, len(tc.input))
		copy(inputCopy, tc.input)

		result := additional.Merge(inputCopy)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("✅ Тест '%s': Пройден\n", tc.name)
		} else {
			fmt.Printf("❌ Тест '%s': ОШИБКА!\n   Вход: %v\n   Ожидали: %v\n   Получили: %v\n",
				tc.name, tc.input, tc.expected, result)
		}
	}
}
