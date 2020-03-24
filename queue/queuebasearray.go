package queue

import "fmt"

/*
* 用数组实现一个顺序队列
 */

// ArrayQueue 基于数组的顺序队列
type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

// NewArrayQueue 构造一个ArrayQueue
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// EnQueue 入队
func (q *ArrayQueue) EnQueue(v interface{}) bool {
	if q.tail == q.capacity {
		return false
	}
	q.q[q.tail] = v
	q.tail++
	return true
}

// DeQueue 出队
func (q *ArrayQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	v := q.q[q.head]
	q.head++
	return v
}

// String 转为字符串
func (q *ArrayQueue) String() string {
	if q.head == q.tail {
		return "empty queue"
	}
	result := "head"
	for i := q.head; i <= q.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", q.q[i])
	}
	result += "<-tail"
	return result
}
