package binarrysearch

import "testing"

func TestBinarySearchRecursive(t *testing.T) {
	a := []int{1, 3, 5, 6, 8, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearchRecursive(a, 17); ret != 9 {
		t.Errorf("test TestBinarySearchRecursive fail, wanted: %v return: %v", 17, ret)
	}

	if ret := BinarySearchRecursive(a, 4); ret != -1 {
		t.Errorf("test TestBinarySearchRecursive fail, wanted: %v return: %v", -1, ret)
	}
}

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearch(a, 17); ret != 9 {
		t.Errorf("test BinarySearch fail, wanted: %v return: %v", 9, ret)
	}

	if ret := BinarySearch(a, 4); ret != -1 {
		t.Errorf("test BinarySearch fail, wanted: %v return: %v", -1, ret)
	}
}

func TestBinarySearchFirst(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 10, 10, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearchFirst(a, 10); ret != 6 {
		t.Errorf("test TestBinarySearchFirst fail, wanted: %v return: %v", 6, ret)
	}

}

func TestBinarySearchLast(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 10, 10, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearchLast(a, 10); ret != 8 {
		t.Errorf("test TestBinarySearchLast fail, wanted: %v return: %v", 8, ret)
	}

}

func TestBinarySearchFirstGT(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 10, 10, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearchFirstGT(a, 10); ret != 9 {
		t.Errorf("test BinarySearchFirstGT fail, wanted: %v return: %v", 9, ret)
	}

}

func TestBinarySearchLastLT(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 8, 10, 10, 10, 11, 13, 15, 17, 18, 20}

	if ret := BinarySearchLastLT(a, 10); ret != 5 {
		t.Errorf("test TestBinarySearchLastLT fail, wanted: %v return: %v", 5, ret)
	}

}
