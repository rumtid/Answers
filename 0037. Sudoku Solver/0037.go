package main

func solveSudoku(board [][]byte) {
	solve(board, 0, 0)
}

func solve(board [][]byte, x, y int) bool {
	n := x*9 + y + 1
	if board[x][y] != '.' {
		return n == 81 || solve(board, n/9, n%9)
	}
	for i := byte('1'); i <= byte('9'); i++ {
		for j := 0; j < 9; j++ {
			if board[x][j] == i || board[j][y] == i ||
				board[x/3*3+j/3][y/3*3+j%3] == i {
				goto L0
			}
		}
		board[x][y] = i
		if n == 81 || solve(board, n/9, n%9) {
			return true
		}
		board[x][y] = byte('.')
	L0:
	}
	return false
}
