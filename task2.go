package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значение квадратов чисел,
	взятых из массива (2,4,6,8,10), и выведет их квадраты в stdout.
*/

func printSquare(x int, wg *sync.WaitGroup) {
	squareX := x * x
	fmt.Println(squareX)
	wg.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printSquare(arr[i], &wg)
	}

	wg.Wait()
}
