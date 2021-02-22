package 盛最多水的容器

import "testing"

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := MaxArea(height)
	if res != 49 {
		t.Error("结果不符合")
		return
	}
	t.Log("success")
}
