package 较大分组的位置

// largeGroupPositions 较大分组的位置
func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0)
	if len(s) <= 2 {
		return res
	}
	// 遍历字符串
	count := 1
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || s[i] != s[i+1] {
			if count >= 3 {
				res = append(res, []int{i - count + 1, i})
			}
			count = 1
		} else {
			count++
		}
	}
	/*for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if count > 3 {
				res = append(res, []int{i - count + 1, i})
			}
			count = 1
		} else {
			count++
		}
	}*/
	return res
}
