package recursion

// Factorial 递归实现阶乘
type Factorial struct {
	val map[int]int
}

// NewFactorial 构造 Factorial
func NewFactorial(n int) *Factorial {
	return &Factorial{
		val: make(map[int]int, n),
	}
}

// Factorial 生成阶乘数据
func (f *Factorial) Factorial(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}

	if n <= 1 {
		f.val[n] = 1
		return f.val[n]
	}

	res := n * f.Factorial(n-1)
	f.val[n] = res
	return res
}

// Print 打印n阶乘的结果
func (f *Factorial) Print(n int) {
	println(f.val[n])
}
