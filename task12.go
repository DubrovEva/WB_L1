package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

	Пояснение к решению:
	По умолчанию в Go не реализовано множество, но можно использовать структуру map[string]struct{}{},
	поскольку сама по себе struct{}{} занимает 0 байт (есть некоторые накладные расходы на ее объявление и хранение
	указателя, но они пренебрежимо малы).
*/

func main() {
	arr := []string{
		"cat",
		"cat",
		"dog",
		"cat",
		"tree",
	}

	stringSet := map[string]struct{}{}
	for _, value := range arr {
		stringSet[value] = struct{}{}
	}

	fmt.Println(stringSet)

}
