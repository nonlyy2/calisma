package leetcode

// Задача №36: Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/

func IsValidSudoku(board [][]byte) bool {
	rows := make([][9]bool, 9)
	cols := make([][9]bool, 9)
	boxes := make([][9]bool, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '1')
			box_index := (i/3)*3 + j/3

			if rows[i][num] {
				return false
			}
			if cols[j][num] {
				return false
			}
			if boxes[box_index][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[box_index][num] = true
		}
	}

	return true
}
