package main

import (
	"log"
	"os"
	"runtime/pprof"

	"github.com/schnell18/testbootcamp/golang/fib"
)

func main() {
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	//nolint:errcheck
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	fibDPLoop()

}

func fibDPLoop() {
	for n := 0; n < 100; n++ {
		seqs := []uint64{12, 14, 30}
		for _, seq := range seqs {
			// fib.FibonacciDynamicProgramming(seq)
			//nolint:staticcheck
			fib.FibonacciRecursive(seq)
		}
	}
}
