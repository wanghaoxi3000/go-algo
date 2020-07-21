package sort

/**
 * @Time   : 2020/7/21 15:00
 * @Author : wanghaoxi3000
 * @Email  : wanghaoxi3000@163.com
 * @Description : 归并排序
 * @Revise : ---
 */

/*MergeSort 归并排序
*	如果要排序一个数组，我们先把数组从中间分成前后两部分，
*	然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。
*
* 递推公式：
* merge_sort(p…r) = merge(merge_sort(p…q), merge_sort(q+1…r))
*
* 终止条件：
* p >= r 不用再继续分解
 */
func MergeSort(arr []int) {
	length := len(arr)

	if len(arr) <= 1 {
		return
	}

	mergeSort(arr, 0, length-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmpArr)
}
