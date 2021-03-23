package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

// lengthOfLongestSubstring 无重复的最长字串
func lengthOfLongestSubstring(s string) int {
	// 存储移动过程中数据
	cache := make(map[byte]int, 0)
	right, res := -1, 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			// i 向前移动了一位,删除之前的元素
			delete(cache, s[i-1])
		}
		// right指针开始右移,同时判断字母是否出现过
		for right+1 < len(s) && cache[s[right+1]] == 0 {
			cache[s[right+1]]++
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
