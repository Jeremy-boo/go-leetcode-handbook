package 最小栈

// MinStack 最小栈
type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor 初始化
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push 向栈中添加元素
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.minStack) == 0 || x <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, x)
	} else {
		s.minStack = append(s.minStack, s.minStack[len(s.minStack)-1])
	}
}

// Pop 删除栈顶元素
func (s *MinStack) Pop() {
	if len(s.stack) < 1 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

// Top 获取栈顶元素
func (s *MinStack) Top() int {
	if len(s.stack) < 1 {
		return -1
	}
	return s.stack[len(s.stack)-1]
}

// GetMin 获取栈里面最小元素
func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
