package 斐波那契数列

// 使用递归的方法 肯定超时
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 优化递归的方法，将每步的结果都缓存起来
func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	f0, f1, f2 := 0, 1, 1
	for i := 2; i <= n; i++ {
		// 根据题意，当N足够大的时候 int32 或者 int64会发生越界
		f2 = (f1 + f0) % 1000000007
		f0 = f1
		f1 = f2
	}
	return f2
}
