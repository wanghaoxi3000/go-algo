package stack

import "fmt"

/*
* 编程模拟实现一个浏览器的前进、后退功能
 */

// Browser 带前进和后退的简易浏览器
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

// NewBrowser 构造一个Browser
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
	}
}

// CanForward 是否可以前进
func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

// CanBack 是否可以后退
func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

// Open 打开一个新的网址
func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	b.forwardStack.Flush()
}

// PushBack 跳转到新网址
func (b *Browser) PushBack(addr string) {
	b.backStack.Push(addr)
}

// Forward 前进
func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

// Back 后退
func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
