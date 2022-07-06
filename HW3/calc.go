package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	c := big.NewInt(1)

	fmt.Println("Калькулятор v0.1")
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Println("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Println("Введите операцию: +,- *, /, ^, !")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
		fmt.Println("Результат выполнения сложения: ", res)
	case "-":
		res = a - b
		fmt.Println("Результат выполнения вычитания: ", res)
	case "*":
		res = a * b
		fmt.Println("Результат выполнения умножения: ", res)
	case "/":
		if b != 0 {
			res = a / b
		} else {
			fmt.Println("Ошибка деления на 0")
			os.Exit(1)
		}
		fmt.Printf("Результат выполнения деления: %.2f\n", res)
	case "^":
		res = math.Pow(a, b)
		fmt.Println("Результат возведения в степень: ", res)
	case "!":
		if a >= 0 {
			c = c.MulRange(1, int64(a))
		} else {
			fmt.Println("Невозможно посчитать факториал отрицательного числа")
			os.Exit(1)
		}
		fmt.Println("Результат вычисления факториала: ", c)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
}
