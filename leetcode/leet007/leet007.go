package leet007

import "math"

// MaxArea 盛水最多的容器
func MaxArea(height []int) int {
	limitLength := int(3 * math.Pow10(4))
	if len(height) < 2 || len(height) > limitLength {
		return 0
	}
	i, j, res := 0, len(height)-1, 0
	for i != j {
		if height[i] < height[j] {
			res = max(res, (j-i)*min(height[i], height[j]))
			i++
		} else {
			res = max(res, (j-i)*min(height[i], height[j]))
			j--
		}
	}
	return res
}

// min math.min
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// max math.max
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
