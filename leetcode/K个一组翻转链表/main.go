package K个一组翻转链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup K个一组翻转链表
// 可采用递归的办法
func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		// 找到下个k数组的起点
		tail = tail.Next
	}
	subResult := reverseKGroup(tail, k)
	// 翻转当前K个节点
	for i := 0; i < k; i++ {
		temp := head.Next
		// 当前节点下个节点指向subResult
		head.Next = subResult
		// 移动当前节点
		subResult = head
		head = temp
	}
	return subResult
}
