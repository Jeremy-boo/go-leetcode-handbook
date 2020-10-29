package leet003


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

// ReverseList 链表反转
func ReverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		prev,cur,cur.Next = cur,cur.Next,prev
	}
	return prev
}