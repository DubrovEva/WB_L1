package main

import (
	"log"
	"reflect"
)

/*
	Разработать программу, которая в рантайме способна определить тип переменной:
	int, string, bool, channel из переменной типа interface{}.

	Пояснение к решению:
	getType -- решение с помощью стандартной библиотеки
	getType2 -- еще более самостоятельное решение
*/

func getType(value interface{}) reflect.Kind {
	return reflect.TypeOf(value).Kind()
}

func getType2(value interface{}) reflect.Kind {
	switch value.(type) {
	case int:
		return reflect.Int
	case string:
		return reflect.String
	case bool:
		return reflect.Bool
	case chan int, chan bool, chan string:
		return reflect.Chan
	}
	return reflect.Invalid
}

func main() {
	tests := map[interface{}]reflect.Kind{
		10:                reflect.Int,
		"10":              reflect.String,
		true:              reflect.Bool,
		make(chan int):    reflect.Chan,
		make(chan string): reflect.Chan,
		make(chan bool):   reflect.Chan,
	}

	for data, ans := range tests {
		result := getType(data)
		if result != ans {
			log.Fatalf("Ошибка на тесте: %v. Ответ: %v, найденный ответ: %v", data, ans, result)
		}
	}

	log.Println("Тесты пройдены")

}
