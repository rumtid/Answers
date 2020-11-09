package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var digits [3][9]bool
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if digits[0][board[i][j]-'1'] {
					return false
				}
				digits[0][board[i][j]-'1'] = true
			}
			if board[j][i] != '.' {
				if digits[1][board[j][i]-'1'] {
					return false
				}
				digits[1][board[j][i]-'1'] = true
			}
			if board[i/3*3+j/3][i%3*3+j%3] != '.' {
				if digits[2][board[i/3*3+j/3][i%3*3+j%3]-'1'] {
					return false
				}
				digits[2][board[i/3*3+j/3][i%3*3+j%3]-'1'] = true
			}
		}
	}
	return true
}
