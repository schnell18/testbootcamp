package fib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciIterative(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		arg  uint64
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2},
		{4, 3}, {5, 5}, {6, 8}, {7, 13},
		{8, 21}, {9, 34}, {10, 55}, {11, 89},
		{12, 144}, {13, 233}, {14, 377},
	}
	if got := FibonacciIterative(0); got != 0 {
		t.Errorf("fibonacci(0) == %d != %d", got, 0)
	}
	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Fibonacci(%d)=%d", tc.arg, tc.want), func(t *testing.T) {
			t.Parallel()
			assert.Equal(tc.want, FibonacciIterative(tc.arg))
		})
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

// Test dynamic programming implementation 2 as sub tests
func TestFibonacciDynamicProgramming2(t *testing.T) {
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
			if got := FibonacciDynamicProgramming2(tc.arg); tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}
