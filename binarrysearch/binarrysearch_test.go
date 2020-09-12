package binarrysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearch(a, 17); ret != 9 {
		t.Errorf("test BinarySearch fail, wanted: %v return: %v", 8, ret)
	}

	if ret := BinarySearch(a, 4); ret != -1 {
		t.Errorf("test BinarySearch fail, wanted: %v return: %v", -1, ret)
	}
}
