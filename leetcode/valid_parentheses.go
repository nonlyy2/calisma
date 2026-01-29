package leetcode

// Задача №20: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

func IsValid(s string) bool {
	var stack []rune
	ch1 := '('
	ch2 := ')'
	ch3 := '['
	ch4 := ']'
	ch5 := '{'
	ch6 := '}'

	for _, ch := range s {
		if ch == ch1 || ch == ch3 || ch == ch5 {
			stack = append(stack, ch)
			continue
		}

		if (ch == ch2 || ch == ch4 || ch == ch6) && len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]
		if top == ch1 && ch == ch2 {
			stack = stack[:len(stack)-1]
			continue
		} else if top == ch3 && ch == ch4 {
			stack = stack[:len(stack)-1]
			continue
		} else if top == ch5 && ch == ch6 {
			stack = stack[:len(stack)-1]
			continue
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
