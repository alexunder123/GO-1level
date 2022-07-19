package sort

// import (
// 	"fmt"
// 	"reflect"
// )

func sort(slice []int) []int {
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

	return slice
}

// func main() {
// 	slice := []int{5, 1, 8, 9, 6, 3, 92, 65, 34, 4, 27, 15}
// 	value := []int{1, 3, 4, 5, 6, 8, 9, 15, 27, 34, 65, 92}
// 	result := sort(slice)
// 	for i, v := range result {
// 		if v != value[i] {
// 			fmt.Println("Массив №1 отсортирован неверно")
// 		}
// 	}
// 	// if reflect.DeepEqual(result, value) == false {
// 	// 	fmt.Println("Массив №1 отсортирован неверно")
// 	// }
// 	fmt.Println(result)
// }
