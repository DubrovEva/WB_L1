package main

import "fmt"

/*
	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}

	Пояснение к решению:
	Тут есть две проблемы.
	Проблема 1. Такой код не сработает на строках с unicode-символами (например, ü может поделиться пополам).
	Проблема 2. slice будет ссылаться на исходную строку, поэтому исходная огромная строка тоже будет храниться
	в памяти.
*/

var justString string

func getPartOfString(runesCount int, text string) string {
	// такое решение будет решать проблему 1, потому что foreach итерируется по рунам, а не по символам,
	// и проблему 2, поскольку ans -- новая строка
	ans := ""
	for i, value := range text {
		if i > runesCount {
			break
		}
		ans = ans + string(value)
	}

	return ans
}

func main() {
	text := "üüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüüü"
	fmt.Println(getPartOfString(11, text))
}
