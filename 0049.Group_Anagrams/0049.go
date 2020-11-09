package leetcode

import "sort"

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, str := range strs {
		bs := ByteSlice(str)
		sort.Sort(bs)
		key := string(bs)
		hash[key] = append(hash[key], str)
	}
	sets := [][]string{}
	for _, set := range hash {
		sets = append(sets, set)
	}
	return sets
}
