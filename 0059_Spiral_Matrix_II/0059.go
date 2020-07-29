package main

func generateMatrix(n int) [][]int {
	var matrix [][]int
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	e := 1
	ilo, ihi := 0, n-1
	jlo, jhi := 0, n-1
	for {
		for k := jlo; k <= jhi; k++ {
			matrix[ilo][k] = e
			e++
		}
		if ilo++; ilo > ihi {
			break
		}
		for k := ilo; k <= ihi; k++ {
			matrix[k][jhi] = e
			e++
		}
		if jhi--; jhi < jlo {
			break
		}
		for k := jhi; k >= jlo; k-- {
			matrix[ihi][k] = e
			e++
		}
		if ihi--; ihi < ilo {
			break
		}
		for k := ihi; k >= ilo; k-- {
			matrix[k][jlo] = e
			e++
		}
		if jlo++; jlo > jhi {
			break
		}
	}
	return matrix
}
