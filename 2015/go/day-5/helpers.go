package main

import (
	"strings"
)

func strSlices(str string) []string {
	strSlices := strings.Split(str, "\n")
	strSlices = strSlices[:len(strSlices)-1] // Remove the last line as it is always empty
	return strSlices
}

func hasThreeVowels(str string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	var counter int
	for _, v := range vowels {
		count := strings.Count(str, v)
		if count > 0 {
			counter += count
		}
	}
	return counter >= 3
}

func hasDoubleLetter(str string) bool {
	var lastChar string
	for _, v := range strings.Split(str, "") {
		if v == lastChar {
			return true
		}
		lastChar = v
	}
	return false
}

func doesNotContainInvalidStrings(str string) bool {
	invalid := []string{"ab", "cd", "pq", "xy"}
	for _, st := range invalid {
		if strings.Contains(str, st) {
			return false
		}
	}
	return true
}

func containsPairOfLetters(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if strings.Count(str, str[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func containsTwoLettersWithOneLetterInbetween(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if byte(str[i]) == str[i+2] {
			return true
		}
	}
	return false
}
