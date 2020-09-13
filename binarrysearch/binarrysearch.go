package binarrysearch

import (
	visulizer "algo/visualizer"
	"fmt"
)

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
