package fibonacci

import "testing"

type test_val struct {
	inp int
	exp int
}

func TestFib_calc(t *testing.T) {
	a, b, d := 0, 1, 2
	tests := []test_val{
		{10, 34},
		{15, 377},
		{22, 10946},
	}
	for _, ts := range tests {
		result := fib_calc(a, b, ts.inp, d)
		if result != ts.exp {
			t.Errorf("стандартная функция. для %d ожидали %d, а получили %d", ts.inp, ts.exp, result)
		}
	}
}

func TestFib_myMap(t *testing.T) {
	tests := []test_val{
		{10, 34},
		{15, 377},
		{22, 10946},
	}
	for _, ts := range tests {
		result := fib_myMap(map[string]int{"prev": 0, "curr": 1, "num": 2, "amount": ts.inp})
		if result != ts.exp {
			t.Errorf("мап функция. для %d ожидали %d, а получили %d", ts.inp, ts.exp, result)
		}
	}
}

func TestFib_map(t *testing.T) {
	tests := []test_val{
		{10, 34},
		{15, 377},
		{22, 10946},
	}
	for _, ts := range tests {
		result := fib_map(ts.inp, map[int]int{1: 0, 2: 1})
		if result != ts.exp {
			t.Errorf("мап функция. для %d ожидали %d, а получили %d", ts.inp, ts.exp, result)
		}
	}
}

func BenchmarkFib_calc(b *testing.B) {
	a, e, d := 0, 1, 2
	c := 20
	for i := 0; i < b.N; i++ {
		fib_calc(a, e, c, d)
	}
}

func BenchmarkFib_myMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_myMap(map[string]int{"prev": 0, "curr": 1, "num": 2, "amount": 20})
	}
}

func BenchmarkFib_map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_map(20, map[int]int{1: 0, 2: 1})
	}
}
