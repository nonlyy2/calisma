package leetcode

// Задача №49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/

// p.s. решение времязатратное, O(nˆ2), попробую перерешать через [26]int

func GroupAnagramsNotOptimal(strs []string) [][]string {
	var res [][]string

	isChecked := make(map[int]bool)
	isChecked[0] = false

	for index, word := range strs {
		freq := make(map[rune]int)

		// не чекать слова которые являются анаграммой другого слова
		if isChecked[index] == true {
			continue
		}

		res = append(res, []string{word})

		// добавляем в чармапу частоту каждой буквы в слове
		for _, ch := range word {
			freq[ch]++
		}

		// проходимся по каждому слову
		for i := index + 1; i < len(strs); i++ {
			buff := make(map[rune]int)
			for k, v := range freq {
				buff[k] = v
			}
			for _, char := range strs[i] {
				buff[char]--
			}

			// ищем анаграммы
			isClearFreq := true
			for _, count := range buff {
				if count != 0 {
					isClearFreq = false
					break
				}
			}

			// если не совпало
			if isClearFreq == false {
				continue
			}
			isChecked[i] = true
			res[len(res)-1] = append(res[len(res)-1], strs[i])
		}
	}
	return res
}
