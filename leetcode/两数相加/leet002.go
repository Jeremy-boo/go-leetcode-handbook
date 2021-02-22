package 两数相加

import "errors"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	var stack1 []interface{}
	var stack2 []interface{}
	// 取出链表中的值放入队列中
	for l1 != nil {
		stack1 = push(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = push(stack2, l2.Val)
		l2 = l2.Next
	}
	// 取出队列中的值进行相加
	carry := 0
	getSumVal := func(x, y int) int {
		value := x + y + carry
		// 发生进位
		if value > 9 {
			value = value - 10
			carry = 1
		} else {
			carry = 0
		}
		return value
	}

	// 插入链表
	var sumL *ListNode
	insert := func(val int) {
		if sumL == nil {
			sumL = &ListNode{
				Val:  val,
				Next: nil,
			}
			return
		}
		node := &ListNode{
			Val:  val,
			Next: sumL,
		}
		sumL = node
	}
	for {
		v1, err1 := pop(&stack1)
		v2, err2 := pop(&stack2)
		if err1 != nil && err2 != nil {
			break
		} else if err1 == nil && err2 == nil {
			insert(getSumVal(v1.(int), v2.(int)))
		} else if err1 == nil {
			insert(getSumVal(v1.(int), 0))
		} else {
			insert(getSumVal(0, v2.(int)))
		}
	}
	return reverseList(sumL)
}

// pop 出栈
func pop(slice *[]interface{}) (interface{}, error) {
	if len(*slice) <= 0 {
		return nil, errors.New("slice is empty")
	}
	value := (*slice)[0]
	*slice = (*slice)[1:]
	return value, nil
}

// push 入栈
func push(slice []interface{}, val interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice)+1, len(slice)+1)
	newSlice[0] = val
	copy(newSlice[1:], slice)
	return newSlice
}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}
