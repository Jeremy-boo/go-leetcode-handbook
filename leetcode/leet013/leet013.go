package leet013

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return dividePow(x, n)
	}
	return 1.0 / dividePow(x, n)
}

func dividePow(r float64, n int) float64 {
	// 终止条件
	if n == 0 {
		return 1.0
	}
	// 处理当前成逻辑
	y := dividePow(r, n/2)
	// 下探到下一层
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * r
	}
}
