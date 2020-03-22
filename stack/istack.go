package stack

// Stack 栈接口
type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush()
}
