package stack

// ArrayStack 基于数组实现的栈
type ArrayStack struct {
	data []interface{} // 数据
	top  int           // 栈顶指针
}

// NewArrayStack 构造一个ArrayStack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// IsEmpty 判断栈是否为空
func (s *ArrayStack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

// Push 数据入栈
func (s *ArrayStack) Push(v interface{}) {
	s.top++
	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

// Pop 数据出栈
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

// Top 返回栈顶数据
func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

// ToIntArray 转换为整型数组
func (s *ArrayStack) ToIntArray() []int {
	array := make([]int, 0)
	for i := 0; i < s.top+1; i++ {
		array = append(array, s.data[i].(int))
	}
	return array
}
