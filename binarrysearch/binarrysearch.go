package binarrysearch

import (
	visulizer "algo/visualizer"
	"fmt"
)

func bs(a []int, v int, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if a[mid] == v {
		return mid
	} else if a[mid] > v {
		return bs(a, v, low, mid-1)
	} else {
		return bs(a, v, mid+1, high)
	}
}

//BinarySearchRecursive 递归二分查找法
func BinarySearchRecursive(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	return bs(a, v, 0, n-1)
}

/*BinarySearch 二分查找法非递归实现
时间复杂度: O(logn)
局限性:
1. 依赖顺序表结构
2. 需要是有序数据
3. 数据量太小或太大都不适合二分查找
*/
func BinarySearch(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	var visual string
	_ = visual

	graph := visulizer.NewGridFromSliceInt(a, v)
	low := graph.AddMarker("low", 0, 0)
	high := graph.AddMarker("high", 0, n-1)
	mid := graph.AddMarker("mid", 0, 0)

	for low.Column <= high.Column {
		mid.Column = (low.Column + high.Column) / 2
		mid.Label = fmt.Sprintf("mid:%d", mid.Column)
		visual = graph.Visulize()

		if a[mid.Column] == v {
			visual = graph.Visulize()
			return mid.Column
		} else if a[mid.Column] > v {
			high.Column = mid.Column - 1
		} else {
			low.Column = mid.Column + 1
		}

		visual = graph.Visulize()
	}

	return -1
}

//BinarySearchFirst 二分查找第一个等于给定值的元素
func BinarySearchFirst(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	var visual string
	_ = visual

	graph := visulizer.NewGridFromSliceInt(a, v)
	low := graph.AddMarker("low", 0, 0)
	high := graph.AddMarker("high", 0, n-1)
	mid := graph.AddMarker("mid", 0, 0)

	for low.Column <= high.Column {
		visual = graph.Visulize()
		mid.Column = (low.Column + high.Column) / 2
		mid.Label = fmt.Sprintf("mid:%d", mid.Column)
		visual = graph.Visulize()

		if a[mid.Column] > v {
			high.Column = mid.Column - 1
		} else if a[mid.Column] < v {
			low.Column = mid.Column + 1
		} else {
			if mid.Column == 0 || a[mid.Column-1] != v {
				visual = graph.Visulize()
				return mid.Column
			}

			high.Column = mid.Column - 1
		}

	}

	return -1
}

//BinarySearchLast 二分查找最后一个值等于给定值的元素
func BinarySearchLast(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	var visual string
	_ = visual

	graph := visulizer.NewGridFromSliceInt(a, v)
	low := graph.AddMarker("low", 0, 0)
	high := graph.AddMarker("high", 0, n-1)
	mid := graph.AddMarker("mid", 0, 0)

	for low.Column <= high.Column {
		visual = graph.Visulize()
		mid.Column = (low.Column + high.Column) / 2
		mid.Label = fmt.Sprintf("mid:%d", mid.Column)
		visual = graph.Visulize()

		if a[mid.Column] > v {
			high.Column = mid.Column - 1
		} else if a[mid.Column] < v {
			low.Column = mid.Column + 1
		} else {
			if mid.Column == n-1 || a[mid.Column+1] != v {
				visual = graph.Visulize()
				return mid.Column
			}

			low.Column = mid.Column + 1
		}

	}

	return -1
}

//BinarySearchFirstGT 二分查找第一个大于等于给定值的元素
func BinarySearchFirstGT(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	var visual string
	_ = visual

	graph := visulizer.NewGridFromSliceInt(a, v)
	low := graph.AddMarker("low", 0, 0)
	high := graph.AddMarker("high", 0, n-1)
	mid := graph.AddMarker("mid", 0, 0)

	for low.Column <= high.Column {
		visual = graph.Visulize()
		mid.Column = (low.Column + high.Column) / 2
		mid.Label = fmt.Sprintf("mid:%d", mid.Column)
		visual = graph.Visulize()

		if a[mid.Column] > v {
			high.Column = mid.Column - 1
		} else if a[mid.Column] < v {
			low.Column = mid.Column + 1
		} else {
			if mid.Column != n-1 && a[mid.Column+1] > v {
				visual = graph.Visulize()
				return mid.Column + 1
			}

			low.Column = mid.Column + 1
		}

	}

	return -1
}

//BinarySearchLastLT 二分查找最后一个小于等于给定值的元素
func BinarySearchLastLT(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	var visual string
	_ = visual

	graph := visulizer.NewGridFromSliceInt(a, v)
	low := graph.AddMarker("low", 0, 0)
	high := graph.AddMarker("high", 0, n-1)
	mid := graph.AddMarker("mid", 0, 0)

	for low.Column <= high.Column {
		visual = graph.Visulize()
		mid.Column = (low.Column + high.Column) / 2
		mid.Label = fmt.Sprintf("mid:%d", mid.Column)
		visual = graph.Visulize()

		if a[mid.Column] > v {
			high.Column = mid.Column - 1
		} else if a[mid.Column] < v {
			low.Column = mid.Column + 1
		} else {
			if mid.Column != n-1 && a[mid.Column-1] < v {
				visual = graph.Visulize()
				return mid.Column - 1
			}

			high.Column = mid.Column - 1
		}

	}

	return -1
}
