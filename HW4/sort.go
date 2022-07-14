package main

import "fmt"

func main() {

	slice := []int{5, 1, 8, 9, 6, 3, 92, 65, 34, 4, 27, 15}

	fmt.Println("Программа сортировки вставкой")
	//fmt.Println("Введите значения через пробел: ")

	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
				j--
			} else {
				break
			}
		}
	}

	fmt.Println("Отсортированные значения: ", slice)
}
