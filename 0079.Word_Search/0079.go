package leetcode

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if search(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[0]) ||
		board[i][j] != word[0] {
		return false
	}
	board[i][j] = '.'
	found := search(board, word[1:], i+1, j) ||
		search(board, word[1:], i-1, j) ||
		search(board, word[1:], i, j+1) ||
		search(board, word[1:], i, j-1)
	board[i][j] = word[0]
	return found
}
