package fib

func FibonacciIterative(n uint64) uint64 {
	var i, x, y uint64 = 0, 0, 1
	for i = 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func FibonacciDynamicProgramming2(n uint64) uint64 {
	memo := make([]uint64, n+1)
	var fibDP func(n uint64) uint64
	fibDP = func(n uint64) uint64 {
		if n < 2 {
			return n
		}
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = fibDP(n-1) + fibDP(n-2)
		return memo[n]
	}
	return fibDP(n)
}

func FibonacciDynamicProgramming(n uint64) uint64 {
	memo := make([]uint64, n+1)
	var fibDP func(n uint64) uint64
	fibDP = func(n uint64) uint64 {
		if n < 2 {
			return n
		}
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = fibDP(n-2) + fibDP(n-1)
		return memo[n]
	}
	return fibDP(n)
}

func FibonacciRecursive(n uint64) uint64 {
	if n <= 2 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciTailRecursive(n uint64) uint64 {
	var fib func(i, j, k uint64) uint64
	fib = func(i, j, k uint64) uint64 {
		if i == n {
			return j + k
		} else {
			return fib(i+1, k, j+k)
		}
	}

	if n < 2 {
		return n
	}
	return fib(2, 0, 1)
}
