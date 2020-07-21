package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
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

	MergeSort(randomList)
	if !reflect.DeepEqual(randomList, sortedList) {
		t.Errorf("test merge sort fail, sort ret: %v", randomList)
	}
	t.Log("sorted list: ", randomList)
}
