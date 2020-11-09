package leetcode

var T = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	s1 := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		s := T[digits[i]-'2']
		s0 := make([]string, 0)
		if i == 0 {
			for j := 0; j < len(s); j++ {
				s0 = append(s0, string(s[j]))
			}
		} else {
			for j := 0; j < len(s1); j++ {
				for k := 0; k < len(s); k++ {
					s0 = append(s0, s1[j]+string(s[k]))
				}
			}
		}
		s1 = s0
	}
	return s1
}
