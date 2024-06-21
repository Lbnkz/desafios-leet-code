package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		m1[word1[i]]++
		m2[word2[i]]++
	}
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m2[k] == 0 {
			return false
		}
	}
	s1 := make([]int, 0, len(m1))
	s2 := make([]int, 0, len(m2))
	for _, v := range m1 {
		s1 = append(s1, v)
	}
	for _, v := range m2 {
		s2 = append(s2, v)
	}
	sort.Ints(s1)
	sort.Ints(s2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	println(closeStrings("abc", "bca")) // true
	println(closeStrings("a", "aa"))     // false
	println(closeStrings("cabbba", "abbccc")) // true
}