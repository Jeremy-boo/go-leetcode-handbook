package leet008

import "testing"

func TestIsValid(t *testing.T) {
	s := "){"
	x := IsValid(s)
	if x != false {
		t.Error("失败!")
		return
	}
	t.Log("Success")
}
