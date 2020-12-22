package leet005

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	nums1 := MoveZeros(nums)
	if ok := reflect.DeepEqual(nums, nums1); !ok {
		t.Error("移动0操作失败")
		return
	}
	t.Log("success")
}

func TestMoveZeros1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	nums1 := MoveZeros1(nums)
	fmt.Println(nums1)
	if ok := reflect.DeepEqual(nums, nums1); !ok {
		t.Error("移动0操作失败")
		return
	}
	t.Log("success")
}
