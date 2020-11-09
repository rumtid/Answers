package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var list []int
	ilo, ihi := 0, len(matrix)-1
	jlo, jhi := 0, len(matrix[0])-1
	for {
		for k := jlo; k <= jhi; k++ {
			list = append(list, matrix[ilo][k])
		}
		if ilo++; ilo > ihi {
			break
		}
		for k := ilo; k <= ihi; k++ {
			list = append(list, matrix[k][jhi])
		}
		if jhi--; jhi < jlo {
			break
		}
		for k := jhi; k >= jlo; k-- {
			list = append(list, matrix[ihi][k])
		}
		if ihi--; ihi < ilo {
			break
		}
		for k := ihi; k >= ilo; k-- {
			list = append(list, matrix[k][jlo])
		}
		if jlo++; jlo > jhi {
			break
		}
	}
	return list
}
