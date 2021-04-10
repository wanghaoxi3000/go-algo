package binarytree

// TreeNode 二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTraversal 前序遍历: 对于树中的任意节点来说，先打印这个节点，然后再打印它的左子树，最后打印它的右子树。
// 前序遍历的递推公式：
// preOrder(r) = print r->preOrder(r->left)->preOrder(r->right)
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Val)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}

	return res
}

// InOrderTraversal 中序遍历: 对于树中的任意节点来说，先打印它的左子树，然后再打印它本身，最后打印它的右子树。
// 中序遍历的递推公式：
// inOrder(r) = inOrder(r->left)->print r->inOrder(r->right)
func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := InOrderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, InOrderTraversal(root.Right)...)

	return res
}

// PostOrderTraversal 后序遍历: 对于树中的任意节点来说，先打印它的左子树，然后再打印它的右子树，最后打印这个节点本身。
// 后序遍历的递推公式：
// postOrder(r) = postOrder(r->left)->postOrder(r->right)->print r
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		lres := PostOrderTraversal(root.Left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.Right != nil {
		rres := PostOrderTraversal(root.Right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.Val)
	return res
}
