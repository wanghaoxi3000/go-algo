package sort

/**
 * @Time   : 2020/7/2 16:59
 * @Author : wanghaoxi3000
 * @Email  : wanghaoxi3000@163.com
 * @Description : 冒泡排序算法
 * @Revise : ---
 */

/*
冒泡排序
1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
3. 针对所有的元素重复以上的步骤，除了最后一个。

平均复杂度: O(n^2)
最好情况: O(n)
最坏情况: O(n^2)
空间复杂度: O(1)
排序方式: In-place
稳定性: 稳定
*/
func BubbleSort(data []int) {
	if len(data) <= 1 {
		return
	}

	// 提前退出标志
	didSwap := false

	length := len(data)
	for i := length; i > 1; i-- {
		didSwap = false

		for j := 0; j < i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				didSwap = true
			}
		}
		if !didSwap {
			break
		}
	}

	return
}
