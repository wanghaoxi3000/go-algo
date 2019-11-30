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
