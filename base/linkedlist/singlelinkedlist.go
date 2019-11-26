package linkedlist

/*
 链表基本数据结构定义
*/

// ListNode 链表节点数据结构
type ListNode struct {
	Next  *ListNode
	Value interface{}
}

// NewListNode 创建新节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}
