package linkedlist

/*
	单链表字符串判断是否为回文字符串
	思路：找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
	时间复杂度：O(N)
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	var prev *ListNode = nil
	var temp *ListNode = nil

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// 反转前半段链表
		temp = slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	// fast = nil 偶数 fast.Next = nil 奇数
	if fast != nil {
		slow = slow.Next
	}

	goFront := prev
	goEnd := slow
	ret := false
	for goFront != nil && goEnd != nil {
		if goEnd.Value != goFront.Value {
			break
		}
		goEnd = goEnd.Next
		goFront = goFront.Next
	}

	ret = (goFront == nil && goEnd == nil)

	// 还原链表
	var next *ListNode
	for prev != nil {
		next = prev.Next
		prev.Next = temp
		temp = prev
		prev = next
	}

	return ret
}
