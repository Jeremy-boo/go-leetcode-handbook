package main

// LRUCache 最近最少使用缓存
type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DoubleLinkedList
}

// DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	value      int
	prev, next *DoubleLinkedList
}

// initDoubleLinkedList 初始化链表
func initDoubleLinkedList(value int) *DoubleLinkedList {
	return &DoubleLinkedList{
		value: value,
	}
}

// Constructor 构造方法
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DoubleLinkedList{},
		size:     0,
	}
	return l
}

// Get 获取相关key
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	//node := this.cache[key]
	return 1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		// 不存在
		node := initDoubleLinkedList(value)
		this.cache[key] = node
		this.size++
		if this.size > this.capacity {

		}
	}
}

func (this *LRUCache) addToHead(node *DoubleLinkedList) {

}
