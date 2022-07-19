package fibonacci

import "fmt"

func fib_calc(a, b, c, d int) int {
	if c == d {
		return b
	}
	d++
	return fib_calc(b, a+b, c, d)
}

func fib_myMap(mapFibo map[string]int) int {
	if mapFibo["num"] == mapFibo["amount"] {
		return mapFibo["curr"]
	}
	mapFibo["prev"], mapFibo["curr"] = mapFibo["curr"], mapFibo["prev"]+mapFibo["curr"]
	mapFibo["num"]++
	return fib_myMap(mapFibo)
}

func fib_map(n int, hash map[int]int) int {
	if val, ok := hash[n]; ok {
		return val
	}
	hash[n] = fib_map(n-1, hash) + fib_map(n-2, hash)
	return hash[n]
}

func main() {
	var a, b, c int
	a, b = 0, 1
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
		d := 2
		fmt.Println("Значение числа Фибоначчи равно ", fib_calc(a, b, c, d))
	}

	hash := map[int]int{
		1: 0,
		2: 1,
	}
	fmt.Println("Значение числа Фибоначчи равно ", fib_map(c, hash))
	fmt.Println("Значение числа Фибоначчи равно ", fib_myMap(mapFibo))
}
