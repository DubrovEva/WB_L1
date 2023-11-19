package main

import (
	"context"
	"log"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
func goroutine1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("горутина 1 остановлена")
			return
		default:
			log.Println("запущена горутина 1")
		}
	}
}

func goroutine2(signalChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-signalChan:
			log.Println("горутина 2 остановлена")
			return
		default:
			log.Println("запущена горутина 2")
		}
	}
}

func main() {
	// wg используется только для проверки того, что все горутины корректно завершили работу
	wg := sync.WaitGroup{}

	// Вариант 1: создаем контекст вместе с функцией cancel, пробрасываем контекст в функцию
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go goroutine1(ctx, &wg)
	time.Sleep(10 * time.Microsecond)
	cancel()

	// Вариант 2: создаем канал, в который попадет сигнал об остановке, пробрасываем его в функцию
	signalChan := make(chan bool)
	wg.Add(1)
	go goroutine2(signalChan, &wg)
	time.Sleep(10 * time.Microsecond)
	signalChan <- true

	wg.Wait()
}
