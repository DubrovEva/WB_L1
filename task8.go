package main

import (
	"log"
)

/*
	Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

	Пояснение к решению: биты считаем с младших разрядов. Двоичные числа в комментариях показаны упрощенно
*/

func changeBit(x int64, i int, toZero bool) int64 {
	// shift -- это двоичное число, у которого на i-ой позиции единица, а на всех остальные -- нули
	shift := int64(1) << (i - 1)
	// bitI -- значение i-го бита ы исходном числе
	bitI := (x & shift) >> (i - 1)

	if toZero && bitI == 1 {
		// чтобы установить на i-ом бите вместо единицы ноль, используется операция сброс бита
		x = x &^ shift
	} else if !toZero && bitI == 0 {
		// чтобы установить на i-ом бите вместо нуля единицу, используется операция дизъюнкция
		x = x | shift
	}

	return x
}

func main() {
	// 1. 10 (b1010), поменять 2-ой бит на 0, ответ 8 (b1000)
	x, ans, i := int64(10), int64(8), 2
	toZero := true
	result := changeBit(x, i, toZero)
	if result != ans {
		log.Fatalf("Неправильный ответ на тесте 1: %v вместо %v", result, ans)
	}

	// 2. 1000 (b1111101000), поменять 6-ой бит на 0, ответ 968 (b1111001000)
	x, ans, i = int64(1000), int64(968), 6
	toZero = true
	result = changeBit(x, i, toZero)
	if result != ans {
		log.Fatalf("Неправильный ответ на тесте 2: %v вместо %v", result, ans)
	}

	// 3. 123 (b1111011), поменять 3-ий бит на 0, число не изменится
	x, ans, i = int64(123), int64(123), 3
	toZero = true
	result = changeBit(x, i, toZero)
	if result != ans {
		log.Fatalf("Неправильный ответ на тесте 3: %v вместо %v", result, ans)
	}

	// 4. 4356 (b1000100000100), поменять 9-ий бит на 1, число не изменится
	x, ans, i = int64(4356), int64(4356), 9
	toZero = false
	result = changeBit(x, i, toZero)
	if result != ans {
		log.Fatalf("Неправильный ответ на тесте 4: %v вместо %v", result, ans)
	}

	// 5. 0 (b0), поменять 9-ий бит на 1, ответ 256 (b100000000)
	x, ans, i = int64(0), int64(256), 9
	toZero = false
	result = changeBit(x, i, toZero)
	if result != ans {
		log.Fatalf("Неправильный ответ на тесте 4: %v вместо %v", result, ans)
	}

	log.Println("Все тесты пройдены")

}
