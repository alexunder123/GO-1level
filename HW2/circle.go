package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64

	fmt.Println("Программа расчета диаметра и длины окружности круга")
	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&s)

	fmt.Println("Диаметр круга равен: ", math.Sqrt(4*s/math.Pi))

	fmt.Println("Длина окружности круга равна: ", 2*math.Pi*math.Sqrt(s/math.Pi))
}
