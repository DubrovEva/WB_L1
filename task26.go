package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные
	(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
*/

func testUniquenessCharacters(s string) bool {
	ss := strings.ToLower(s)

	isLetterInString := make(map[int32]bool)
	for _, letter := range ss {
		if isLetterInString[letter] {
			return false
		}
		isLetterInString[letter] = true
	}

	return true
}

func main() {
	arr := []string{
		"",          // true
		"a",         // true
		"aa",        // false
		"aA",        // false
		"aB",        // true
		"abcd",      // true
		"abCdefAaf", // false
		"aabcd",     // false
	}

	for _, value := range arr {
		fmt.Println(testUniquenessCharacters(value))
	}
}
