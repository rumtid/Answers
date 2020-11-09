package leetcode

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	total := 0
	table := make(map[string]int)
	for _, word := range words {
		table[word] += 1
		total += 1
	}

	m := 0
	n := len(words[0])
	r := make([]int, 0)
	g := make(map[string]int)
	for i, j := 0, 0; j < len(s)-n+1; {
		c, ok := table[s[j:j+n]]
		if !ok || g[s[j:j+n]] == c {
			m, g = 0, make(map[string]int)
			i, j = i+1, i+1
			continue
		}
		m += 1
		g[s[j:j+n]] += 1
		if m == total {
			r = append(r, i)
			m, g = 0, make(map[string]int)
			i, j = i+1, i+1
			continue
		}
		j += n
	}

	return r
}
