package binarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	pre, in, post []int
}{

	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{3, 2, 1},
	},

	{
		[]int{1, 2, 5, 5, 3, 6, 7},
		[]int{4, 2, 5, 1, 6, 3, 7},
		[]int{4, 5, 2, 6, 7, 3, 1},
	},
}

func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

func Test_PreOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.pre, PreOrderTraversal(root), "输入:%v", tc)
	}
}

func Test_InOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.in, InOrderTraversal(root), "输入:%v", tc)
	}
}

func Test_PostOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.post, PostOrderTraversal(root), "输入:%v", tc)
	}
}
