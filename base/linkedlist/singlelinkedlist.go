package linkedlist

import "fmt"

/*
 单链表基本数据结构定义
*/

// ListNode 链表节点数据结构
type ListNode struct {
	Next  *ListNode
	Value interface{}
}

// LinkedList 链表头节点
type LinkedList struct {
	Head   *ListNode
	Length uint
}

// NewListNode 创建新节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// NewListNodeFromArray 从数组生成链表
func NewListNodeFromArray(array []interface{}) *LinkedList {
	head := &LinkedList{Head: &ListNode{}}
	cur := head.Head

	for _, v := range array {
		cur.Next = &ListNode{Value: v}
		head.Length++
		cur = cur.Next
	}

	return head
}

// ToIntArray 将链表转换为整型数组
func (h *LinkedList) ToIntArray() []int {
	array := make([]int, h.Length)
	cur := h.Head.Next

	for i := 0; i < int(h.Length); i++ {
		array[i] = cur.Value.(int)
		cur = cur.Next
	}

	return array
}

// Print 打印链表
func (h *LinkedList) Print() {
	cur := h.Head.Next
	format := fmt.Sprintf("Length: %v Value: ", h.Length)
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Value)
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

/*
Reverse 反转链表
时间复杂度：O(N)
*/
func (h *LinkedList) Reverse() {
	if h.Head == nil || h.Head.Next == nil || h.Head.Next.Next == nil {
		return
	}

	cur := h.Head
	var preNode, nextNode *ListNode
	for cur != nil {
		nextNode = cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}

	h.Head.Next = preNode
}
