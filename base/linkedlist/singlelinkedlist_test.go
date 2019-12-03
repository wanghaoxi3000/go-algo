package linkedlist

import (
	"reflect"
	"testing"
)

var sourceData []interface{}
var l *LinkedList

func init() {
	sourceData = []interface{}{1, 2, 3, 4, 5}
	l = NewListNodeFromArray(sourceData)
}

func TestReverse(t *testing.T) {
	testList := NewListNodeFromArray(sourceData)
	testList.Reverse()
	exceptList := []int{5, 4, 3, 2, 1}
	result := testList.ToIntArray()

	if !reflect.DeepEqual(exceptList, result) {
		t.Errorf("Test function Reverse() fail, except %v result %v", exceptList, result)
	}
}

func TestHasCycle(t *testing.T) {
	testList := NewListNodeFromArray(sourceData)
	if testList.HasCycle() {
		t.Errorf("Test function HasCycle() fail, except false result true")
	}

	testList.Head.Next.Next.Next.Next.Next.Next = testList.Head.Next.Next.Next
	if !testList.HasCycle() {
		t.Errorf("Test function HasCycle() fail, except true result false")
	}
}
