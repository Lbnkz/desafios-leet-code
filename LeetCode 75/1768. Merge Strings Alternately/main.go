package main

func mergeAlternately(w1 string, w2 string) string {
	var r string
	for i := 0; i < len(w1) || i < len(w2); i++ {
		if i < len(w1) {
			r += string(w1[i])
		}
		if i < len(w2) {
			r += string(w2[i])
		}
	}
	return r
}


func main() {
	println(mergeAlternately("abc", "pqr")) // "apbqcr"
	println(mergeAlternately("ab", "pqrs")) // "apbqrs"
	println(mergeAlternately("abcd", "pq")) // "apbqcd"
}