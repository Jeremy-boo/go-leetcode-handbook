package 验证二叉搜索树

import "testing"

func TestIsValidBST(t *testing.T) {
	n1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	n2 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	n3 := TreeNode{
		Val:   4,
		Left:  &n1,
		Right: &n2,
	}
	ok := isValidBST(&n3)
	if !ok {
		t.Log("success")
	}
}
