package 反转链表2

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
	cacheNode := &ListNode{
		Next: head,
	}
	preM := cacheNode
	// 找到m位置前节点
	for i := 1; i <= m-1; i++ {
		preM = preM.Next
	}
	// 反转m节点到n节点之间的链表
	var linkedList *ListNode
	cur := preM.Next
	for j := m; j <= n; j++ {
		linkedList, cur, cur.Next = cur, cur.Next, linkedList
	}

	// preM.Next为m节点，原始m节点应该指向n+1节点
	preM.Next.Next = cur
	preM.Next = linkedList

	return cacheNode.Next
}
