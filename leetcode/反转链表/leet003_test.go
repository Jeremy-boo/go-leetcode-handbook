package 反转链表

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	result := ReverseList(&l1)
	fmt.Println(result)
}

func TestReverseList2(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	result := ReverseList2(&l1)
	fmt.Println(result)
}
