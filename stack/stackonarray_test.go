package stack

import (
	"reflect"
	"testing"
)

func TestArrayStackPush(t *testing.T) {
	s := NewArrayStack()
	expectData := []int{1, 2, 3}
	for _, v := range expectData {
		s.Push(v)
	}

	stackArray := s.ToIntArray()
	if !reflect.DeepEqual(expectData, stackArray) {
		t.Errorf("Test ArrayStack pop error, expect: %v, return: %v", expectData, stackArray)
	}
}

func TestArrayStackPop(t *testing.T) {
	s := NewArrayStack()
	testData := []int{1, 2, 3}
	for _, v := range testData {
		s.Push(v)
	}

	for range testData {
		s.Pop()
	}
	s.Pop()

	expectData := []int{}
	stackArray := s.ToIntArray()
	if !reflect.DeepEqual(expectData, stackArray) {
		t.Errorf("Test ArrayStack pop error, expect: %v, return: %v", expectData, stackArray)
	}
}

func TestArrayStackTop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 2; i > 0; i-- {
		s.Pop()
		v := s.Top()
		if v != i {
			t.Errorf("Test ArrayStack top error, expect: %v, return: %v", i, v)
		}
	}

	s.Pop()
	v := s.Top()
	if v != nil {
		t.Errorf("Test ArrayStack top error, expect: %v, return: %v", nil, v)
	}
}
