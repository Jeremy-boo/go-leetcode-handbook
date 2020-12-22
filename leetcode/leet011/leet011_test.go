package leet011

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
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
	n4 := invertTree(&n3)
	fmt.Println(n4)
}
