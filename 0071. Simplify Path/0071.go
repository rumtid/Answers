package main

import "strings"

func simplifyPath(path string) string {
	var dirs []string
	strs := strings.Split(path, "/")
	for _, str := range strs {
		switch str {
		case "", ".":
			continue
		case "..":
			if n := len(dirs); n > 0 {
				dirs = dirs[:n-1]
			}
		default:
			dirs = append(dirs, str)
		}
	}
	return "/" + strings.Join(dirs, "/")
}
