package 多数元素

import "testing"

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(nums)
	if res != 2 {
		t.Error("失败")
		return
	}
	t.Log("Pass")
}
