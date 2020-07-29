package main

import "regexp"

var re, _ = regexp.Compile("\\A *[\\+-]?(\\d+(\\.\\d*)?|\\.\\d+)(e[\\+-]?\\d+)? *\\z")

func isNumber(s string) bool {
	return re.MatchString(s)
}
