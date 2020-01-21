package linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	strs := map[string]bool{
		"": true,
		"heooeh": true,
		"hello": false,
		"heoeh": true,
		"a": true,
	}

	for str, expect := range strs {
		h := NewListNode(nil)	
		cur := h
		if len(str) > 0 {
			cur.Value = int32(str[0])// 下标截取的类型是 uint8
		}
		if len(str) > 1{
			for _, ch := range str[1:] {
				cur.Next = NewListNode(ch)
				cur = cur.Next
			}
		}
		
		ret := isPalindrome(h)
		if expect != ret {
			t.Error("string:", str, "ret is:", ret, "expect is:", expect)
		}
	}
}
