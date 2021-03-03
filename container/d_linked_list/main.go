package d_linked_list

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

// Insert 添加到任意位置
func (dl *DList) Insert(value interface{}, index int) {

}
