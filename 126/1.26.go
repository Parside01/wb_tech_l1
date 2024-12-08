package main

func IsUniqChars(input string) bool {
	m := make(map[rune]int)
	for _, v := range input {
		m[v]++

		if m[v] == 2 {
			return false
		}
	}
	return true
}
