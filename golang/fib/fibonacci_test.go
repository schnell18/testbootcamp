package fib

import (
	"fmt"
	"testing"
)

func TestFibonacciIterative(t *testing.T) {
	if got := FibonacciIterative(0); got != 0 {
		t.Errorf("fibonacci(0) == %d != %d", got, 0)
	}
	if got := FibonacciIterative(1); got != 1 {
		t.Errorf("fibonacci(1) == %d != %d", got, 1)
	}
	if got := FibonacciIterative(2); got != 1 {
		t.Errorf("fibonacci(2) == %d != %d", got, 1)
	}
	if got := FibonacciIterative(3); got != 2 {
		t.Errorf("fibonacci(3) == %d != %d", got, 2)
	}
	if got := FibonacciIterative(4); got != 3 {
		t.Errorf("fibonacci(4) == %d != %d", got, 3)
	}
	// ...
	if got := FibonacciIterative(14); got != 377 {
		t.Errorf("fibonacci(14) == %d != %d", got, 377)
	}
}

// Test recursive implementation as sub tests
func TestFibonacciRecursive(t *testing.T) {
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
			if got := FibonacciRecursive(tc.arg); tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

// Test dynamic programming implementation as sub tests
func TestFibonacciDynamicProgramming(t *testing.T) {
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
			if got := FibonacciDynamicProgramming(tc.arg); tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

// Test dynamic programming implementation as sub tests
func TestFibonacciTailRecursive(t *testing.T) {
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
			if got := FibonacciTailRecursive(tc.arg); tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}
