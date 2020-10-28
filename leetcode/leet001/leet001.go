package leet001

import (
	"errors"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 链表
type ListNode struct {
	Val int
	Next *ListNode
}

// AddTwoNumbers 两数之和2
func AddTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode {
	// 取出第一个数的数据
	var stack1 []interface{}
	for l1 != nil {
		stack1 = push(stack1,l1.Val)
		l1 = l1.Next
	}

	var stack2 []interface{}
	for l2 != nil {
		stack2 = push(stack2,l2.Val)
		l2 = l2.Next
	}

	var sumL *ListNode
	flag := 0
	// 插入链表
	insert := func(x int) {
		if sumL == nil {
			sumL = &ListNode{
				Val:  x,
			}
			return
		}
		node := &ListNode{
			Val: x ,
			Next: sumL,
		}
		sumL = node
	}
	// 各个位数相加，判断进位
	getSum := func(v1,v2 int) int{
		value := v1+v2+flag
		if value > 9 {
			value = value -10
			flag =1
		}else {
			flag =0
		}
		return value
	}
	for {
		v1,err1 := pop(&stack1)
		v2,err2 := pop(&stack2)
		if err1 != nil && err2 != nil {
			break
		}else if err1 == nil && err2 == nil {
			insert(getSum(v1.(int),v2.(int)))
		}else if err1 == nil  {
			insert(getSum(v1.(int),0))
		}else  {
			insert(getSum(0,v2.(int)))
		}
	}
	if flag ==1 {
		insert(1)
	}
	return sumL
}

func pop(slice *[]interface{}) (value interface{},err error) {
	if len(*slice)== 0{
		return nil,errors.New("stack is empty")
	}
	value = (*slice)[0]
	*slice = (*slice)[1:]
	return value,nil
}

func push(slice []interface{},value interface{}) []interface{} {
	newSlice := make([]interface{},len(slice)+1,len(slice)+1)
	newSlice[0] = value
	copy(newSlice[1:],slice)
	return newSlice
}


func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre,cur,cur.Next = cur,cur.Next,pre
	}
	return pre
}