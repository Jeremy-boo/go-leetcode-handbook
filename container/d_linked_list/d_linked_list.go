package d_linked_list

import "errors"

// DNode 链表节点
type DNode struct {
	data       interface{}
	prev, next *DNode
}

// DList 双端链表
type DList struct {
	size       int
	head, tail *DNode
}

// Constructor 双端链表初始化
func Constructor() DList {
	return DList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// Size 获取链表的长度
func (dl *DList) Size() int {
	return dl.size
}

// AddHeadNode 添加数据到头节点
func (dl *DList) AddHeadNode(value interface{}) {
	node := &DNode{
		data: value,
		prev: nil,
		next: nil,
	}
	if dl.Size() == 0 {
		dl.head = node
		dl.tail = node
	} else {
		dl.head.prev = node
		node.prev = nil
		node.next = dl.head
		dl.head = node
	}
	dl.size++
}

// Append 添加数据去尾节点
func (dl *DList) Append(value interface{}) {
	node := &DNode{
		data: value,
		prev: nil,
		next: nil,
	}
	if dl.Size() == 0 {
		dl.head = node
		dl.tail = node
	} else {
		dl.tail.next = node
		node.prev = dl.tail
		node.next = nil
		dl.tail = node
	}
	dl.size++
}

// Head 头节点
func (dl *DList) Head() *DNode {
	return dl.head
}

// Tail 尾节点
func (dl *DList) Tail() *DNode {
	return dl.tail
}

// InsertByIndex 添加到任意位置
func (dl *DList) InsertByIndex(value interface{}, index int) (bool, error) {
	if index > dl.Size() && index < 0 {
		return false, errors.New("index is not available")
	}
	if index == 0 || dl.Size() == 0 {
		dl.AddHeadNode(value)
		return true, nil
	}
	if index == dl.Size() {
		dl.Append(value)
		return true, nil
	}
	cur := dl.head
	node := &DNode{
		data: value,
	}
	for i := 1; i < dl.Size(); i++ {
		if i == index {
			node.next = cur.next
			cur.next.prev = node
			cur.next = node
			node.prev = cur
			dl.size++
			break
		}
		cur = cur.next
	}
	return true, nil
}

func (dl *DList) FindByIndex(index int) (*DNode, error) {
	if index < 0 || index > dl.Size()-1 {
		return nil, errors.New("index is not available")
	}
	if index == 0 {
		return dl.head, nil
	}
	if index == dl.Size()-1 {
		return dl.tail, nil
	}
	cur := dl.head
	for i := 1; i < dl.Size(); i++ {
		cur = cur.next
		if i == index {
			break
		}
	}
	return cur, nil
}
