package 两数之和

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoSum1(nums, target)
	fmt.Println(res)
	targetRes := []int{0, 1}
	if !reflect.DeepEqual(res, targetRes) {
		t.Error("结果错误")
		return
	}
	t.Log("success")
}
