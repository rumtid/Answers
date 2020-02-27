package main

var solutions int
var col, diag1, diag2 []bool

func totalNQueens(n int) int {
	solutions = 0
	col = make([]bool, n)
	diag1 = make([]bool, n*2-1)
	diag2 = make([]bool, n*2-1)
	solve(n, 0)
	return solutions
}

func solve(n, i int) {
	if i == n {
		solutions++
		return
	}
	for j, c := range col {
		if c || diag1[i+j] || diag2[i-j+n-1] {
			continue
		}
		col[j] = true
		diag1[i+j] = true
		diag2[i-j+n-1] = true
		solve(n, i+1)
		diag2[i-j+n-1] = false
		diag1[i+j] = false
		col[j] = false
	}
}
