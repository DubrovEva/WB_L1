package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает,
	вычитает две числовых переменных a,b, значение которых > 2^20.

	Пояснение к решению:
	В Go базово встроена длинная арифметика -- в пакете Big.
*/

func main() {
	a, b := big.NewInt(int64(math.Pow(2, 30))), big.NewInt(int64(math.Pow(2, 25)))
	var result big.Int
	// Сложение
	fmt.Println(result.Add(a, b)) // 1107296256

	// Вычитание
	fmt.Println(result.Sub(a, b)) // 1040187392

	// Умножение
	fmt.Println(result.Mul(a, b)) // 36028797018963968

	// Деление
	fmt.Println(result.Div(a, b)) // 32

}
