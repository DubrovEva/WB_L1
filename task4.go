package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

	Пояснение к решению: воркеры завершают работу с помощью контекста (он прокидывается во все функции, и когда
	канал SignalChan ловит сигнал Interrupt, запускается CancelFunc контекста).
*/

func producer(ctx context.Context, wg *sync.WaitGroup, mainChan chan<- int) {
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

func worker(ctx context.Context, wg *sync.WaitGroup, mainChan <-chan int, workerID int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case value := <-mainChan:
			fmt.Printf("Worker#%d прочитал сообщение %d\n", workerID, value)
			waitMs := rand.Intn(500) + 500
			time.Sleep(time.Duration(waitMs) * time.Millisecond)
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	mainChan := make(chan int, 5)

	var workerCount int
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, mainChan, i)
	}

	wg.Add(1)
	go producer(ctx, &wg, mainChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	cancel()
	wg.Wait()
	fmt.Printf("\nПрограмма была остановлена\n")
}
