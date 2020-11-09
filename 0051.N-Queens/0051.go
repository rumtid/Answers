package leetcode

import "bytes"

var solution []string
var solutions [][]string
var col, diag1, diag2 []bool

func solveNQueens(n int) [][]string {
	solution = []string{}
	solutions = [][]string{}
	col = make([]bool, n)
	diag1 = make([]bool, n*2-1)
	diag2 = make([]bool, n*2-1)
	solve(n, 0)
	return solutions
}

func solve(n, i int) {
	if i == n {
		solutions = append(solutions, append([]string{}, solution...))
		return
	}
	str := bytes.Repeat([]byte("."), n)
	for j, c := range col {
		if c || diag1[i+j] || diag2[i-j+n-1] {
			continue
		}
		str[j] = 'Q'
		col[j] = true
		diag1[i+j] = true
		diag2[i-j+n-1] = true
		solution = append(solution, string(str))
		solve(n, i+1)
		solution = solution[:len(solution)-1]
		diag2[i-j+n-1] = false
		diag1[i+j] = false
		col[j] = false
		str[j] = '.'
	}
}
