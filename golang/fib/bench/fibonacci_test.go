package bench_test

import (
	"testing"

	"github.com/schnell18/testbootcamp/golang/fib"
)

var result uint64

func BenchmarkFibonacciIterative(b *testing.B) {
	seqs := []uint64{12, 14, 30}
	var r uint64
	for j := 0; j < b.N; j++ {
		for _, seq := range seqs {
			r = fib.FibonacciIterative(seq)
		}
	}
	result = r
}

func BenchmarkFibonacciDynamicProgramming(b *testing.B) {
	seqs := []uint64{12, 14, 30}
	var r uint64
	for j := 0; j < b.N; j++ {
		for _, seq := range seqs {
			r = fib.FibonacciDynamicProgramming(seq)
		}
	}
	result = r
}

func BenchmarkFibonacciDynamicProgramming2(b *testing.B) {
	seqs := []uint64{12, 14, 30}
	var r uint64
	for j := 0; j < b.N; j++ {
		for _, seq := range seqs {
			r = fib.FibonacciDynamicProgramming2(seq)
		}
	}
	result = r
}

func BenchmarkFibonacciTailRecursive(b *testing.B) {
	seqs := []uint64{12, 14, 30}
	var r uint64
	for j := 0; j < b.N; j++ {
		for _, seq := range seqs {
			r = fib.FibonacciTailRecursive(seq)
		}
	}
	result = r
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	seqs := []uint64{12, 14, 30}
	var r uint64
	for j := 0; j < b.N; j++ {
		for _, seq := range seqs {
			r = fib.FibonacciRecursive(seq)
		}
	}
	result = r
}

// func BenchmarkFibonacciRecursiveForever(b *testing.B) {
// 	for j := 0; j < b.N; j++ {
//         _ = fib.FibonacciRecursive(uint64(j))
// 	}
// }
