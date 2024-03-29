package calci

import (
	"testing"
)

func Test_Add(t *testing.T) {
	sum := Sum(2, 3)
	if sum != 5 {
		t.Error("Expected", 5, "Got", sum)
	}

	type test struct {
		data []int
		sum  int
	}
	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{5, 5, 5}, 15},
		{[]int{-1, 21}, 20},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.sum {
			t.Error("Expected", v.sum, "Got", x)
		}
	}
}
