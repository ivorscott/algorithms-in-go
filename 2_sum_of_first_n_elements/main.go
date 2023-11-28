package main

import "fmt"

/*
	Problem:
	Find the sum of the first N numbers.

	Objective function:
	f(i) is the sum of the first i elements.

	Recurrence relation:
	f(n) = f(n-1) + n

n = 3
1+2+3 = 6
n = 5
1+2+3+4+5 = 15
 */
func main() {
	fmt.Println(nSum(3))
}

func nSum(n int) int {
	// what happens if n is 0?
	x := make([]int, n + 1)
	x[0] = 0
	for i := 1; i < n + 1; i++ {
		x[i] = x[i-1] + i
	}
	return x[n]
}
