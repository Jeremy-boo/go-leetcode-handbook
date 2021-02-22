package 两数之和

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSum1 两数之和
func TwoSum1(nums []int, target int) []int {
	cache := make(map[int]int)
	for index, val := range nums {
		if _, ok := cache[target-val]; ok {
			// 存在
			return []int{cache[target-val], index}
		}
		cache[val] = index
	}
	return nil
}
