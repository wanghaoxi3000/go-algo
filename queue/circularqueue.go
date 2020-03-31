package queue

import "fmt"

// CircularQueue 循环队列
type CircularQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

// NewCircularQueue 构造一个循环数组
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

// IsEmpty 判断队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

// IsFull 判断队列是否已满
func (q *CircularQueue) IsFull() bool {
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

// EnQueue 入队
func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.q[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// DeQueue 出队
func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.q[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.q[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
