package leet006

import "testing"

func TestClimbStairs(t *testing.T) {
	n := 4
	r := ClimbStairs(n)
	if r != 5 {
		t.Error("测试失败！")
		return
	}
	t.Log("Success")
}
