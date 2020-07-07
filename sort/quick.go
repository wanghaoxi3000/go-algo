package sort

/**
 * @Time   : 2020/7/7 14:11
 * @Author : wanghaoxi3000
 * @Email  : wanghaoxi3000@163.com
 * @Description : 快速排序
 * @Revise : ---
 */

/*
快速排序
1. 从数列中挑出一个元素，称为 “基准”（pivot）;
2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；

平均复杂度: O(n log n)
最好情况: O(n log n)
最坏情况: O(n^2)
空间复杂度: O(log n)
排序方式: In-place
稳定性: 不稳定
*/

func QuickSort(data []int) {
	separateSort(data, 0, len(data)-1)
}

// 递归分割函数
func separateSort(arr []int, start, end int) {
	if start > end {
		return
	}

	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

// 分区排序函数
func partition(arr []int, start, end int) int {
	// 选取最后一位对比数字
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return i
}
