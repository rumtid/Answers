package leetcode

var (
	T1 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	T2 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	T3 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	T4 = []string{"", "M", "MM", "MMM"}
)

func intToRoman(num int) string {
	return T4[num/1000] + T3[num/100%10] + T2[num/10%10] + T1[num%10]
}
