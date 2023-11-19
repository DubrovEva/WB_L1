package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.

	Пояснение к решению:
	В Go есть готовая реализация: atomic.Int64, но потенциально можно реализовать через Mutex,
	почему реализация через Mutex ниже.
*/

type Counter struct {
	value int64
}

func (x *Counter) AddOne() {
	atomic.AddInt64(&x.value, 1)
}

func main() {
	wg := sync.WaitGroup{}
	//counter := CounterByMutex{Mutex: &sync.Mutex{}, value: 0}
	counter := Counter{0}

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				counter.AddOne()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

type CounterByMutex struct {
	*sync.Mutex
	value int64
}

func (cbm *CounterByMutex) AddOne() {
	cbm.Lock()
	cbm.value += 1
	cbm.Unlock()
}
