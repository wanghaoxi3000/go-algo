package stack

import "algo/linkedlist"

// LinkedListStack 基于链表实现的栈
type LinkedListStack struct {
	topNode *linkedlist.ListNode //栈顶节点
}

// NewLinkedListStack 构造一个LinkedListStack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

// IsEmpty 判断栈是否为空
func (s *LinkedListStack) IsEmpty() bool {
	return s.topNode == nil
}

// Push 数据入栈
func (s *LinkedListStack) Push(v interface{}) {
	s.topNode = &linkedlist.ListNode{
		Next:  s.topNode,
		Value: v,
	}
}

// Pop 数据出栈
func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.topNode.Value
	s.topNode = s.topNode.Next
	return v
}

// Top 返回栈顶数据
func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.topNode.Value
}

// Flush 清空栈
func (s *LinkedListStack) Flush() {
	s.topNode = nil
}

// ToIntArray 转换为整型数组
func (s *LinkedListStack) ToIntArray() []int {
	array := make([]int, 0)
	ptr := s.topNode
	for ptr != nil {
		array = append(array, ptr.Value.(int))
		ptr = ptr.Next
	}
	return array
}
