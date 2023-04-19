package fib

import (
	"fmt"
	"reflect"
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
			if got := Fibonacci(tc.arg); tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

func TestFibonacci2(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 uint64
	}{
		//TODO: Add test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := Fibonacci2(tArgs.n)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Fibonacci2 got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
