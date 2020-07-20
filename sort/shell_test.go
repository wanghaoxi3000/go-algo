package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

/**
 * @Time   : 2020/7/20 17:10
 * @Author : wanghaoxi3000
 * @Email  : wanghaoxi3000@163.com
 * @Description : 希尔排序测试
 * @Revise : ---
 */

func TestShellSort(t *testing.T) {
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

	ShellSort(randomList)
	if !reflect.DeepEqual(randomList, sortedList) {
		t.Errorf("test shell sort fail, sort ret: %v", randomList)
	}
	t.Log("sorted list: ", randomList)
}
