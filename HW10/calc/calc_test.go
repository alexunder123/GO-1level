package calc

import (
	"fmt"
	"math"
	"testing"
)

type test_val struct {
	a   float64
	b   float64
	op  string
	res float64
}

func TestCalc(t *testing.T) {
	tests := []test_val{
		{10, 34, "*", 340},
		{15.1, 37.7, "+", 52.80},
		{22, 1096, "/", 0.02},
		{242, 196, "-", 46},
		{42, 96.4, "/", 0.44},
	}
	for _, ts := range tests {
		result := calulator(ts.a, ts.b, ts.op)
		if math.Round(result*100)/100 != ts.res {
			t.Errorf("ожидали %.2f, а получили %.2f", ts.res, result)
		}
	}
}

func ExampleCalc() {
	fmt.Println(calulator(5.2, 2, "/"))
	//output 2.6
}
