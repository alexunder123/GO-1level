package calc

import (
	"fmt"
	"math"
	"os"
)

func calulator(a float64, b float64, op string) float64 {
	var res float64

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b != 0 {
			res = a / b
		} else {
			fmt.Println("Ошибка деления на 0")
			os.Exit(1)
		}
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	return res
}
