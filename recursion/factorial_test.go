package recursion

import "testing"

func TestFac_Factorial(t *testing.T) {
	fac := NewFactorial(15)
	for i := 1; i < 15; i++ {
		fac.Factorial(i)
		fac.Print(i)
	}
}
