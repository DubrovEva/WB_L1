package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать паттерн «адаптер» на любом примере.

	Пример: Допустим, у нас в проекте используются и классические словари, и безопасные словари из библиотеки sync.
	Мы не можем поменять сами классы, но нам нужен общий интерфейс для получения элемента из словаря.
*/

type MapAdapter interface {
	GetElem(any) any
}

type SafeMapAdapter struct {
	sync.Map
}

func (m *SafeMapAdapter) GetElem(key any) any {
	result, exist := m.Load(key)
	if exist {
		return result
	} else {
		return nil
	}
}

type UnsafeMapAdapter struct {
	Map map[any]any
}

func (m *UnsafeMapAdapter) GetElem(key any) any {
	result, exist := m.Map[key]
	if exist {
		return result
	} else {
		return nil
	}
}

func main() {
	wild := "wild"
	berries := "berries"
	bear := "bear"

	unsafeMap := map[any]any{wild: berries}
	safeMap := sync.Map{}
	safeMap.Store(wild, bear)

	unsafeMapAdapter := UnsafeMapAdapter{unsafeMap}
	safeMapAdapter := SafeMapAdapter{safeMap}

	mapArray := []MapAdapter{&unsafeMapAdapter, &safeMapAdapter}

	for _, value := range mapArray {
		fmt.Println(value.GetElem(wild))
	}

}
