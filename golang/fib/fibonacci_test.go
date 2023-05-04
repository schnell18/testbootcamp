package fib

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		arg  uint64
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2},
		{4, 3}, {5, 5}, {6, 8}, {7, 13},
		{8, 21}, {9, 34}, {10, 55}, {11, 89},
		{12, 144}, {13, 233}, {14, 377},
	}
	for i, tc := range cases {
		if got := Fibonacci(tc.arg); tc.want != got {
			t.Errorf("TC#%d fibonacci(%d) == %d != %d", i+1, tc.arg, got, tc.want)
		}
	}
}

// Test recursive implementation as sub tests
func TestFibonacciSub(t *testing.T) {
	cases := []struct {
		arg  uint64
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2},
		{4, 3}, {5, 5}, {6, 8}, {7, 13},
		{8, 21}, {9, 34}, {10, 55}, {11, 89},
		{12, 144}, {13, 233}, {14, 377},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Fibonacci(%d)=%d", tc.arg, tc.want), func(t *testing.T) {
			if got := Fibonacci2(tc.arg); tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}
