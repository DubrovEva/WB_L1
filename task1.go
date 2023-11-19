package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s \n", h.name)
}

type Action struct {
	designation string
	Human
}

func (a *Action) SayHi() {
	fmt.Printf("Hi, I am %s, and I do %s", a.name, a.designation)
}

func main() {
	human := Human{"Eva", 20}
	action := Action{"jogging", human}

	human.SayHi()
	action.SayHi()
}
