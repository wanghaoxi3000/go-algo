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
	testList.Print()
	testList.Reverse()
	testList.Print()
	exceptList := []int{5, 4, 3, 2, 1}
	result := testList.ToIntArray()

	if !reflect.DeepEqual(exceptList, result) {
		t.Errorf("Test function Reverse() fail, except %v result %v", exceptList, result)
	}
}

func TestHasCycle(t *testing.T) {
	l.Reverse()
	l.Print()
}
