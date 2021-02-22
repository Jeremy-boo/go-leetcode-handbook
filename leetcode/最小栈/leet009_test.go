package 最小栈

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(-2)
	min1 := obj.GetMin()
	if min1 != -2 {
		t.Error("获取最小栈失败!")
		return
	}
	t.Log("success")

	obj.Pop()
	min2 := obj.GetMin()
	if min2 != 1 {
		t.Error("获取最小栈失败!")
		return
	}
	t.Log("success")
}
