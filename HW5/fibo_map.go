package main

import "fmt"

func fib_calc(mapFibo map[string]int) int {
	if mapFibo["num"] == mapFibo["amount"] {
		return mapFibo["curr"]
	}
	mapFibo["prev"], mapFibo["curr"] = mapFibo["curr"], mapFibo["prev"]+mapFibo["curr"]
	mapFibo["num"]++
	return fib_calc(mapFibo)
}

func main() {
	var c int
	fmt.Println("Введите порядковый номер числа Фибоначчи: ")
	fmt.Scanln(&c)

	mapFibo := map[string]int{
		"prev":   0,
		"curr":   1,
		"num":    2,
		"amount": c,
	}

	if c <= 0 {
		fmt.Println("Введено неверное значение")
	} else if c == 1 {
		fmt.Println("Значение числа Фибоначчи равно 0")
	} else if c == 2 {
		fmt.Println("Значение числа Фибоначчи равно 1")
	} else {
		fmt.Println("Значение числа Фибоначчи равно ", fib_calc(mapFibo))
	}

}
