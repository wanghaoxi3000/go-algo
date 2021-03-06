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

// MergeSortedList 合并有序链表
func MergeSortedList(a, b *LinkedList) *LinkedList {
	if a == nil || a.Head == nil || a.Head.Next == nil {
		return b
	}
	if b == nil || b.Head == nil || b.Head.Next == nil {
		return a
	}

	mergedList := LinkedList{Head: &ListNode{}}
	mergedPtr := mergedList.Head
	aPtr := a.Head.Next
	bPtr := b.Head.Next

	for aPtr != nil && bPtr != nil {
		if aPtr.Value.(int) < bPtr.Value.(int) {
			mergedPtr.Next = aPtr
			aPtr = aPtr.Next
		} else {
			mergedPtr.Next = bPtr
			bPtr = bPtr.Next
		}
		mergedPtr = mergedPtr.Next
	}

	if aPtr != nil {
		mergedPtr.Next = aPtr
	} else {
		mergedPtr.Next = bPtr
	}

	return &mergedList
}

// ToIntArray 将链表转换为整型数组
func (h *LinkedList) ToIntArray() []int {
	array := make([]int, 0)
	cur := h.Head.Next

	for cur != nil {
		array = append(array, cur.Value.(int))
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

	cur := h.Head.Next
	var preNode, nextNode *ListNode
	for cur != nil {
		nextNode = cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}

	h.Head.Next = preNode
}

// HasCycle 检测链表是否有环
func (h *LinkedList) HasCycle() bool {
	if h.Head == nil || h.Head.Next == nil {
		return false
	}

	slow := h.Head.Next
	fast := h.Head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// DeleteBottomNode 删除链表倒数第n个节点
func (h *LinkedList) DeleteBottomNode(n int) {
	if n <= 0 || nil == h.Head || nil == h.Head.Next {
		return
	}

	fast := h.Head
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return
		}
	}

	slow := h.Head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
}

// FindMiddleNode 获取中间节点
func (h *LinkedList) FindMiddleNode() *ListNode {
	if nil == h.Head || nil == h.Head.Next {
		return nil
	}
	if nil == h.Head.Next.Next {
		return h.Head.Next
	}

	slow, fast := h.Head, h.Head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
