package main

import "fmt"

func main() {
	var a, b float32

	fmt.Println("Программа расчета площади прямоугольника")
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Println("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника равна: ", a*b)
}
