package main

import (
	"fmt"
	"math"
)

func main() {
	var number0 int
	var number float64

	fmt.Println("Программа выделения единиц, десятков и сотен из числа")
	fmt.Println("Введите целое трехзначное число: ")
	fmt.Scanln(&number0)

	number = float64(number0)
	fmt.Println("Число единиц равно: ", math.Mod(number, 10))
	number = math.Trunc(number / 10)
	fmt.Println("Число десятков равно: ", math.Mod(number, 10))
	number = math.Trunc(number / 10)
	fmt.Println("Число сотен равно: ", math.Mod(number, 10))
}
