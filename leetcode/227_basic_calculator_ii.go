package leetcode

// Задача №227: Basic Calculator II
// https://leetcode.com/problems/basic-calculator-ii/

func Calculate(s string) int {
	res := 0
	stack := []int{}
	operation := '+'
	currNum := 0

	for i := range s {
		if isDigit(s[i]) {
			currNum = currNum*10 + int(s[i]-'0')
		}

		if (!isDigit(s[i]) && s[i] != ' ') || i == len(s)-1 {
			switch operation {
			case '+':
				stack = append(stack, currNum)

			case '-':
				stack = append(stack, -currNum)

			case '*':
				stack[len(stack)-1] *= currNum

			case '/':
				stack[len(stack)-1] /= currNum
			}

			operation = rune(s[i])
			currNum = 0
		}
	}

	for _, val := range stack {
		res += val
	}

	return res
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
