package sort

/**
 * @Time   : 2020/7/3 14:16
 * @Author : wanghaoxi3000
 * @Email  : wanghaoxi3000@163.com
 * @Description : 选择排序算法
 * @Revise : ---
 */


/*
选择排序
1. 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
2. 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3. 重复第二步，直到所有元素均排序完毕。

平均复杂度: O(n^2)
最好情况: O(n^2)
最坏情况: O(n^2)
空间复杂度: O(1)
排序方式: In-place
稳定性: 不稳定
*/
func SelectionSort(data []int)  {
	length := len(data)

	for minIndex := 0; minIndex < length - 1; minIndex++ {
		for selectIndex := minIndex+1;selectIndex < length; selectIndex++ {
			if data[selectIndex] < data[minIndex] {
				data[selectIndex], data[minIndex] = data[minIndex], data[selectIndex]
			}
		}
	}

	return
}
