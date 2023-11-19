package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
*/

func reverseArrayString(text string) string {
	stringArray := strings.Fields(text)
	n := len(stringArray)
	for i := 0; i < n/2; i += 1 {
		stringArray[i], stringArray[n-1-i] = stringArray[n-1-i], stringArray[i]
	}
	return strings.Join(stringArray, " ")
}

func main() {
	test := "snow dog sun"
	fmt.Println(reverseArrayString(test)) // sun dog snow
}
