package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.

	Пояснение к решению: я явно реализовала безопасный map, но в Go уже есть sync.Map, который решает ту же задачу
*/

type SafeMap struct {
	sync.Mutex
	Map map[interface{}]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{Map: map[interface{}]interface{}{}}
}

func (sm *SafeMap) Set(key interface{}, value interface{}) {
	sm.Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}

func (sm *SafeMap) Get(key interface{}) interface{} {
	sm.Lock()
	defer sm.Unlock()
	return sm.Map[key]
}

func main() {
	safeMap := NewSafeMap()

	safeMap.Set("key", "value")
	fmt.Println(safeMap.Get("key"))
}
