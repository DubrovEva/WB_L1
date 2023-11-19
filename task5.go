package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func sender(ctx context.Context, wg *sync.WaitGroup, mainChan chan<- int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())

	for {
		num := rand.Int()
		select {
		case <-ctx.Done():
			return
		case mainChan <- num:
			fmt.Printf("Producer записал в канал сообщение %d\n", num)
		}
	}
}

func recipient(ctx context.Context, wg *sync.WaitGroup, mainChan <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case value := <-mainChan:
			fmt.Printf("Recipient получил сообщение %d\n", value)
			waitMs := rand.Intn(500) + 500
			time.Sleep(time.Duration(waitMs) * time.Millisecond)
		}
	}
}

const (
	timeToStop = 5 * time.Second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeToStop)
	wg := sync.WaitGroup{}
	mainChan := make(chan int, 5)

	wg.Add(1)
	go recipient(ctx, &wg, mainChan)

	wg.Add(1)
	go sender(ctx, &wg, mainChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	select {
	case <-signalChan:
		cancel()
	case <-ctx.Done():
	}

	wg.Wait()
	fmt.Printf("\nПрограмма была остановлена\n")
}
