package leet006

// ClimbStairs 爬楼梯
// n=1   1
// n=2   2
// n=3   3
// n=4   即 f(n) = f(n-1)+ f(n-2)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i <= n; i++ {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
	}
	return f3
}
