package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива,
	во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func produceDataFromArray(wg *sync.WaitGroup, arr []int, first chan<- int) {
	defer wg.Done()
	for _, value := range arr {
		first <- value
	}
	close(first)
}

func produceDataFromChan(wg *sync.WaitGroup, first <-chan int, second chan<- int) {
	defer wg.Done()
	for value := range first {
		valueDouble := value * 2
		second <- valueDouble
	}
	close(second)
}

func printDataToStdout(wg *sync.WaitGroup, second <-chan int) {
	defer wg.Done()

	for value := range second {
		fmt.Println(value)
	}
}

func main() {
	wg := sync.WaitGroup{}

	arr := rand.Perm(10)
	fmt.Println(arr)
	firstChan, secondChan := make(chan int), make(chan int)

	wg.Add(3)
	go produceDataFromArray(&wg, arr, firstChan)
	go produceDataFromChan(&wg, firstChan, secondChan)
	go printDataToStdout(&wg, secondChan)

	wg.Wait()
}
