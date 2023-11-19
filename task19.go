package main

import (
	"log"
)

/*
	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/

func reverseString(s string) string {
	// не оптимальное решение, потому что каждый раз копируется строка result
	result := ""
	for _, value := range s {
		result = string(value) + result
	}
	return result
}

func reverseString2(s string) string {
	result := []rune(s)
	for i := 0; i < len(result)/2; i += 1 {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return string(result)
}

func main() {
	tests := map[string]string{
		"a":    "a",
		"世界":   "界世",
		"abc":  "cba",
		"你好好好": "好好好你",
	}

	for key, value := range tests {
		result := reverseString(key)
		if result != value {
			log.Fatalf("Не пройден тест. Введенная строка: %q, правильный ответ: %q, найденный ответ: %q", key, value, result)
		}
		result2 := reverseString2(key)
		if result2 != value {
			log.Fatalf("Не пройден тест 2. Введенная строка: %q, правильный ответ: %q, найденный ответ: %q", key, value, result2)
		}

	}

	log.Println("Тесты пройдены")

}
