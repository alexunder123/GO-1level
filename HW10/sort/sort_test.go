package sort

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	slice := []int{5, 1, 8, 9, 6, 3, 92, 65, 34, 4, 27, 15}
	value := []int{1, 3, 4, 5, 6, 88, 9, 15, 27, 34, 65, 92}
	result := sort(slice)
	for i, v := range result {
		if v != value[i] {
			t.Errorf("Массив №1 отсортирован неверно")
		}
	}

	slice = []int{3, 10, 84, 19, 6, 63, 92, 5, 3, 4, 2, 15, 22}
	value = []int{2, 3, 3, 4, 5, 6, 10, 15, 19, 22, 63, 84, 92}
	result = sort(slice)
	if reflect.DeepEqual(result, value) == false {
		t.Errorf("Массив №2 отсортирован неверно")
	}
}

func TestSortTestify(t *testing.T) {
	slice := []int{5, 1, 8, 9, 6, 3, 92, 65, 34, 4, 27, 15}
	value := []int{1, 3, 4, 5, 6, 88, 9, 15, 27, 34, 65, 92}
	assert.Equal(t, sort(slice), value)
}
