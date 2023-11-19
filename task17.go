package main

import (
	"fmt"
	"log"
)

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

var errorMsg = "ошибка на тесте номер %d. Найденный ответ: %v, правильный ответ: %v\n"

func binarySearchInArray(arr []int, x int) int {
	// Поиск элемента в отсортированном массиве

	l := -1
	r := len(arr) - 1
	for r-l > 1 {
		m := (r + l) / 2
		if arr[m] < x {
			l = m
		} else {
			r = m
		}
	}

	if arr[r] != x {
		return -1
	} else {
		return r
	}
}

func binarySearchByAnswer(arr []int, f func(int) bool) int {
	// Бинарный поиск по ответу (поиск первого значения, удовлетворяющему некоторому условию)

	l := -1
	r := len(arr) - 1
	for r-l > 1 {
		m := (r + l) / 2
		if f(arr[m]) {
			r = m
		} else {
			l = m
		}
	}
	if f(arr[r]) {
		return r
	} else {
		return -1
	}
}

func binarySearchByReal(l, r, e float64, f func(float64) float64) float64 {
	// Бинарный поиск по вещественным числам (поиск значения x, при котором функция равна 0)
	// e -- погрешность решения

	for r-l > e {
		m := (r + l) / 2
		if f(m) > 0 {
			r = m
		} else {
			l = m
		}
	}
	return l
}

func testBinarySearchInArray() error {
	// Задача: Найти минимальную позицию, на которой стоит указанный элемент.

	arr := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 9, 10}

	tests := map[int]int{
		0:  -1,
		1:  0,
		5:  4,
		10: 12,
	}

	testNum := 0
	for key, ans := range tests {
		testNum += 1
		result := binarySearchInArray(arr, key)
		if result != ans {
			return fmt.Errorf(errorMsg, testNum, result, ans)
		}
	}

	return nil
}

func testBinarySearchByAnswer() error {
	// Задача: В массиве сначала идут двойки, а потом единицы. Найти минимальную позицию, на которой стоит единица

	testNum := 0
	arr := []int{1, 1, 1, 1, 1, 1, 1}
	ans := 0
	result := binarySearchByAnswer(arr, func(x int) bool { return x == 1 })
	if result != ans {
		return fmt.Errorf(errorMsg, testNum, result, ans)
	}

	testNum = 1
	arr = []int{2, 2, 2, 2, 2, 2, 2}
	ans = -1
	result = binarySearchByAnswer(arr, func(x int) bool { return x == 1 })
	if result != ans {
		return fmt.Errorf(errorMsg, testNum, result, ans)
	}

	testNum = 2
	arr = []int{2, 2, 2, 2, 2, 2, 1}
	ans = 6
	result = binarySearchByAnswer(arr, func(x int) bool { return x == 1 })
	if result != ans {
		return fmt.Errorf(errorMsg, testNum, result, ans)
	}

	testNum = 3
	arr = []int{2, 2, 2, 1, 1, 1, 1}
	ans = 3
	result = binarySearchByAnswer(arr, func(x int) bool { return x == 1 })
	if result != ans {
		return fmt.Errorf(errorMsg, testNum, result, ans)
	}

	return nil
}

func testBinarySearchByReal() error {
	/*
		Задача: Найти кубический корень числа с погрешностью 0.00006
		при границах l = -100, r = 100
	*/
	l, r, e := -100., 100., 0.00006

	tests := map[float64]float64{
		27.:           3,
		125:           5,
		2.82842712475: 1.41421,
	}

	testNum := 0
	for cube, ans := range tests {
		testNum += 1
		result := binarySearchByReal(l, r, e, func(x float64) float64 { return x*x*x - cube })
		if result > ans+e || result < ans-e {
			return fmt.Errorf(errorMsg, testNum, result, ans)
		}
	}

	return nil
}

func main() {
	err := testBinarySearchInArray()
	if err != nil {
		log.Fatalf("Не пройдены тесты на бинпоиск по массиву: %v", err)
	}
	log.Println("Пройдены тесты на бинпоиск по массиву")

	err = testBinarySearchByAnswer()
	if err != nil {
		log.Fatalf("Не пройдены тесты на бинпоиск по ответу: %v", err)
	}
	log.Println("Пройдены тесты на бинпоиск по ответу")

	err = testBinarySearchByReal()
	if err != nil {
		log.Fatalf("Не пройдены тесты на бинпоиск по действительным числам: %v", err)
	}
	log.Println("Пройдены тесты на бинпоиск по действительным числам")
}
