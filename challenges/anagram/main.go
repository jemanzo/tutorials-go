package main

import (
	"fmt"
)

func main() {
	var num uint = 0
	// TRUE
	num++
	fmt.Printf("ANAGRAM %v is %v\n", num, areAnagrams("abc", "cba"))
	num++
	fmt.Printf("ANAGRAM %v is %v\n", num, areAnagrams("Hello, 世界", "世界, Hello"))

	// FALSE
	num++
	fmt.Printf("ANAGRAM %v is %v\n", num, areAnagrams("ab", "ab "))
	num++
	fmt.Printf("ANAGRAM %v is %v\n", num, areAnagrams("abcb", "cbaa"))
	num++
	fmt.Printf("ANAGRAM %v is %v\n", num, areAnagrams("Hello, 世界", "hello, 世界"))
}

// ANAGRAMS
// Two strings are anagrams when they have the same
// amount of chars with the same occurrence frequency

func areAnagrams(s1, s2 string) bool {
	chars1 := []rune(s1)
	chars2 := []rune(s2)

	if len(chars1) != len(chars2) {
		return false
	}

	counterMap1 := mapCharsOf(chars1)
	counterMap2 := mapCharsOf(chars2)
	// fmt.Println(counterMap1)
	// fmt.Println(counterMap2)

	for key := range counterMap1 {
		if counterMap1[key] != counterMap2[key] {
			return false
		}
	}

	return true
}

func mapCharsOf(chars []rune) map[string]int {
	counterMap := make(map[string]int)
	for v := range chars {
		ch := string(chars[v])
		if counterMap[ch] == 0 {
			counterMap[ch] = 1
		} else {
			counterMap[ch]++
		}
	}
	return counterMap
}
