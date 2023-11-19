package main

import (
	"fmt"
	"sort"
)

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножествах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(arr)

	ans := map[int]map[float64]struct{}{}

	for _, value := range arr {
		key := int(value) - int(value)%10
		group, exists := ans[key]
		if !exists {
			group = map[float64]struct{}{}
			ans[key] = group
		}
		group[value] = struct{}{}
	}

	for key, group := range ans {
		fmt.Println(key, group)
	}
}
