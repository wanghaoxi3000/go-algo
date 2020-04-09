package recursion

import "fmt"

/*
* 递归实现斐波那契数列
* F[n]=F[n-1]+F[n-2](n>=3,F[1]=1,F[2]=1)
* 1，1，2，3，5，8
 */

// Fibs 斐波那契数列
type Fibs struct {
	val map[int]int
}

// NewFibs 构造一个斐波那契数据结构R
func NewFibs(n int) *Fibs {
	return &Fibs{
		make(map[int]int, n),
	}
}

// Fibonacci 生成斐波那契
func (fib *Fibs) Fibonacci(n int) int {
	if fib.val[n] != 0 {
		return fib.val[n]
	}

	if n <= 1 {
		fib.val[1] = 1
		return 1
	} else if n == 2 {
		fib.val[2] = 1
		return 1
	}

	res := fib.Fibonacci(n-1) + fib.Fibonacci(n-2)
	fib.val[n] = res
	return res
}

// Print 打印数据
func (fib *Fibs) Print(n int) {
	fmt.Println(fib.val[n])
}
