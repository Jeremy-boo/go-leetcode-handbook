package leet004

import "strconv"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseBetween 按照指定位置反转链表
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	return nil
}
