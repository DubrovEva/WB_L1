package main

import "fmt"

/*
	Удалить i-ый элемент из слайса.
*/

func deleteElem(slice []int, indexToRm int) []int {
	// Решение за линию с сохранением порядка элементов

	result := make([]int, len(slice)-1)
	mark := 0 // сдвиг на 1 в случае, когда мы прошли удаляемый элемент
	for i, _ := range result {
		if i == indexToRm {
			mark = 1
		}
		result[i] = slice[i+mark]
	}

	return result
}

func main() {
	slice := []int{2, 3, 5, 7}

	fmt.Println(deleteElem(slice, 0)) // [3, 5, 7]
	fmt.Println(deleteElem(slice, 1)) // [2, 5, 7]
	fmt.Println(deleteElem(slice, 2)) // [2, 3, 7]
	fmt.Println(deleteElem(slice, 3)) // [2, 3, 5]
}

func deleteElem2(slice []int, indexToRm int) []int {
	// Решение за константу с несохранением порядка элементов

	sliceSize := len(slice)
	slice[indexToRm], slice[sliceSize-1] = slice[sliceSize-1], slice[indexToRm]
	slice = slice[:sliceSize-1]

	return slice
}

func deleteElem3(slice []int, indexToRm int) []int {
	/*
		Интересный пример с StackOverflow: он работает, но только один раз, потому что
		функция append меняет массив, на который указывает исходный слайс.
		(в случае реалокации немного по-другому)

		Если запустить эту функцию на примере из main:
			0. Изначально массив выглядит как [2, 3, 5, 7], указатель ведет на элемент 2, длина 4.
			1. При вызове функции первый раз, в slice[:indexToRm] указатель
			будет ссылаться на первый элемент в массиве [2, 3, 5, 7], и в него append докидает элементы
			из slice[indexToRm+1:] на места старых элементов, то есть указатель будет ссылаться на массив [3, 5, 7, 7].
			Ответ будет правильным, [3, 5, 7].
			2. На втором шаге slice[:indexToRm] равен [3], а slice[indexToRm+1:] -- на [7, 7]. Ответ будет [3, 7, 7],
			указатель исходного слайса будет ссылаться на [3, 7, 7, 7].
			Ну и дальше все плохо :)
	*/

	result := append(slice[:indexToRm], slice[indexToRm+1:]...)
	return result
}
