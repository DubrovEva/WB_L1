package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.

	Пояснение к решению:
	В Go нет классического Set, но в качестве Set можно использовать Map.
	Map в Go реализован как хэш-таблица, поэтому получение заданного элемента будет за константу.
	Таким образом текущее решение работает за линию от длины меньшей мапы, и это оптимальная асимптотика,
	поскольку в любом решении необходимо просмотреть все элементы меньшей мапы.
*/

func intersection(first, second map[interface{}]struct{}) map[interface{}]struct{} {
	result := map[interface{}]struct{}{}

	if len(first) > len(second) {
		first, second = second, first
	}

	for key := range first {
		_, exist := second[key]
		if exist {
			result[key] = struct{}{}
		}
	}

	return result
}

func main() {

	first := map[interface{}]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}

	second := map[interface{}]struct{}{
		1: {},
		2: {},
		6: {},
	}

	fmt.Println(intersection(first, second))

}
