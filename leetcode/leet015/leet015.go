package leet015

func majorityElement(nums []int) int {
	// 第一种暴力破解法
	// 第二种 利用hashmap计算每个元素出现的次数，然后再获取最多次数的key
	cache := make(map[int]int, 0)
	// 得出每个元素的次数
	for _, v := range nums {
		if _, ok := cache[v]; !ok {
			// 不存在
			cache[v] = 1
		} else {
			cache[v] = cache[v] + 1
		}
	}
	tempCount := 0
	var n int
	for key, val := range cache {
		if val > tempCount {
			tempCount = val
			n = key
		}
	}
	return n
}
