package 反转链表

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

// ReverseList 链表反转
func ReverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}

// ReverseList2 链表反转 第2种解法
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var beg *ListNode = nil
	mid := head
	end := head.Next
	for {
		// 修改mid 的向
		mid.Next = beg
		// 当end 为空的时候退出
		if end == nil {
			break
		}
		// 整体向后移动
		beg = mid
		mid = end
		end = end.Next
	}
	head = mid
	return head
}
