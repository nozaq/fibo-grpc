package handlers

import (
	"fmt"
	"testing"
)

var fiboTable = map[int]int{
	1: 1,
	2: 1,
	3: 2,
	4: 3,
	5: 5,
	6: 8,
	7: 13,
}

func TestFibo(t *testing.T) {
	for k, ev := range fiboTable {
		av := fibo(k)
		if ev != av {
			t.Error(fmt.Sprintf("expected fibo(%d) to be %d, but got %d", k, ev, av))
		}
	}
}
