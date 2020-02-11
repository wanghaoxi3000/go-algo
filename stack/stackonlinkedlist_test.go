package stack

import (
	"reflect"
	"testing"
)

func TestLinkedListStackPush(t *testing.T) {
	s := NewLinkedListStack()
	expectData := []int{3, 2, 1}
	for i := range expectData {
		s.Push(i + 1)
	}

	stackArray := s.ToIntArray()
	if !reflect.DeepEqual(expectData, stackArray) {
		t.Errorf("Test LinkedListStack pop error, expect: %v, return: %v", expectData, stackArray)
	}
}

func TestLinkedListStackPop(t *testing.T) {
	s := NewLinkedListStack()
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
		t.Errorf("Test LinkedListStack pop error, expect: %v, return: %v", expectData, stackArray)
	}
}

func TestLinkedListTop(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 2; i > 0; i-- {
		s.Pop()
		v := s.Top()
		if v != i {
			t.Errorf("Test LinkedListStack top error, expect: %v, return: %v", i, v)
		}
	}

	s.Pop()
	v := s.Top()
	if v != nil {
		t.Errorf("Test LinkedListStack top error, expect: %v, return: %v", nil, v)
	}
}
