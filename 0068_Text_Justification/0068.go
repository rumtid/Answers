package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var lines []string
	for len(words) > 0 {
		var i, w int
		for ; i < len(words) && w+len(words[i])+i <= maxWidth; i++ {
			w += len(words[i])
		}

		var line string
		n := maxWidth - w - (i - 1)
		if i == 1 || i == len(words) {
			line += strings.Join(words[:i], " ")
			line += strings.Repeat(" ", n)
		} else {
			for j := 0; j < n; j++ {
				words[j%(i-1)] += " "
			}
			line = strings.Join(words[:i], " ")
		}

		lines = append(lines, line)
		words = words[i:]
	}
	return lines
}
