package bench_test

import (
	"fmt"
	"testing"

	"github.com/schnell18/testbootcamp/golang/fib"
)

var result uint64

func BenchmarkFibonacciIterative(b *testing.B) {
	seqs := []uint64{30, 50, 100}
	var r uint64
	for _, seq := range seqs {
		b.Run(fmt.Sprintf("fib(%d)", seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = fib.FibonacciIterative(seq)
			}
		})
	}
	result = r
}

func BenchmarkFibonacciDynamicProgramming(b *testing.B) {
	seqs := []uint64{30, 50, 100}
	var r uint64
	for _, seq := range seqs {
		b.Run(fmt.Sprintf("fib(%d)", seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = fib.FibonacciDynamicProgramming(seq)
			}
		})
	}
	result = r
}

func BenchmarkFibonacciTailRecursive(b *testing.B) {
	seqs := []uint64{30, 50, 100}
	var r uint64
	for _, seq := range seqs {
		b.Run(fmt.Sprintf("fib(%d)", seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = fib.FibonacciTailRecursive(seq)
			}
		})
	}
	result = r
}

func BenchmarkFibonacciDynamicProgramming2(b *testing.B) {
	seqs := []uint64{30, 50, 100}
	var r uint64
	for _, seq := range seqs {
		b.Run(fmt.Sprintf("fib(%d)", seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = fib.FibonacciDynamicProgramming2(seq)
			}
		})
	}
	result = r
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	// NOTE: use smaller nubmer as recursive algorithm is very slow for big number
	seqs := []uint64{12, 14, 30}
	var r uint64
	for _, seq := range seqs {
		b.Run(fmt.Sprintf("fib(%d)", seq), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = fib.FibonacciRecursive(seq)
			}
		})
	}
	result = r
}

// func BenchmarkFibonacciRecursiveForever(b *testing.B) {
//     for j := 0; j < b.N; j++ {
//         _ = fib.FibonacciRecursive(uint64(j))
//     }
// }
