package main

import (
	"fmt"
	"sync"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func getSquare(x int, result chan int, wg *sync.WaitGroup) {
	squareX := x * x
	result <- squareX
	wg.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	sum := 0
	wg := sync.WaitGroup{}

	result := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go getSquare(arr[i], result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for value := range result {
		sum += value
	}

	fmt.Println(sum)
}
