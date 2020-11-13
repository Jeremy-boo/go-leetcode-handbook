package leet005

// MoveZeros 移动0
// 解法1. 移动非0元素到前面位置，记录补0的开始位置
func MoveZeros(nums []int) []int {
	zeroIndex := 0
	for _, num := range nums {
		if num != 0 {
			// 移动非0元素到前面位置
			nums[zeroIndex] = num
			zeroIndex++
		}
	}
	// 填充0
	for j := zeroIndex; j < len(nums); j++ {
		nums[j] = 0
	}
	return nums
}

// MoveZeros1 移动0第二种解法
func MoveZeros1(nums []int) []int {
	cache := make([]int, 0)
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			cache = append(cache, nums[i])
		} else {
			zeroCount++
		}
	}
	// 补0
	for i := 0; i < zeroCount; i++ {
		cache = append(cache, 0)
	}
	copy(nums, cache)
	return cache
}
