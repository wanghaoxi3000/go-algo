package linkedlist

import (
	"reflect"
	"testing"
)

var sourceData []interface{}

func init() {
	sourceData = []interface{}{1, 2, 3, 4, 5}
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

func TestMergeSortedList(t *testing.T) {
	testList1 := NewListNodeFromArray([]interface{}{1, 3, 5, 7, 9})
	testList2 := NewListNodeFromArray([]interface{}{2, 4, 6, 8, 10})

	exceptList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mergedList := MergeSortedList(testList1, testList2)
	if !reflect.DeepEqual(exceptList, mergedList.ToIntArray()) {
		t.Errorf("Test function Reverse() fail, except %v result %v", exceptList, mergedList.ToIntArray())
	}
}

func TestDeleteBottomNode(t *testing.T) {
	testList := NewListNodeFromArray([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
	testList.DeleteBottomNode(3)

	exceptList := []int{1, 2, 3, 4, 5, 6, 8, 9}
	if !reflect.DeepEqual(exceptList, testList.ToIntArray()) {
		t.Errorf("Test function DeleteBottomNode() fail, except %v result %v", exceptList, testList.ToIntArray())
	}
}
