package fib

func Fibonacci(n uint64) uint64 {
	var i, x, y uint64 = 0, 0, 1
	for i = 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func Fibonacci2(n uint64) uint64 {
	if n <= 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci3(n uint64) uint64 {
    switch n {
    case 0:
        return 0
    case 1:
		return 1
	default:
        return fib(n, 2, 0, 1)
    }
}

func fib(n, i, j, k uint64) uint64 {
    if i == n {
        return j + k
    } else {
        return fib(n, i + 1, k, j+k)
    }
}

