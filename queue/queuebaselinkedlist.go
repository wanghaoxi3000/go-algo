package queue

import "fmt"

/*
* 用链表实现一个链式队列
 */

// ListNode 链式队列节点
type ListNode struct {
	val  interface{}
	next *ListNode
}

// LinkedListQueue 基于数组的链式队列
type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

// NewLinkedListQueue 构造一个
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

// EnQueue 入队
func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if nil == q.tail {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

// DeQueue 出队
func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.val
	q.head = q.head.next
	q.length--
	return v
}

// String 转为字符串
func (q *LinkedListQueue) String() string {
	if q.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
