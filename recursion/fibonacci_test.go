package recursion

import "testing"

func TestFac_Factorial(t *testing.T) {
	fac := NewFibs(10)
	for i := 1; i < 15; i++ {
		fac.Fibonacci(i)
		fac.Print(i)
	}
}
