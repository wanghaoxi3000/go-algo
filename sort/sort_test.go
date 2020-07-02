package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

/**
 * @Time   : 2020/7/2 17:25
 * @Author : wanghaoxi3000
 * @Email  : wanghaoxi3000@163.com
 * @Description : 排序算法测试
 * @Revise : ---
 */

func TestBubbleSort(t *testing.T) {
	randomList := make([]int, 100)
	sortedList := make([]int, 100)

	rand.Seed(time.Now().UnixNano())
	var num int
	for i := range randomList {
		num = rand.Intn(100)
		randomList[i] = num
		sortedList[i] = num
	}
	sort.Ints(sortedList)

	BubbleSort(randomList)
	if !reflect.DeepEqual(randomList, sortedList) {
		t.Errorf("test bubble sort fail, sort ret: %v", randomList)
	}
	t.Log(randomList)
	t.Log(sortedList)
}