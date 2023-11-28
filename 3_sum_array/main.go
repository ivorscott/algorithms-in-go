package main

import "fmt"

/*
	Problem:
	Given an array as an input argument, write a method that will return the sum of all the integer
	elements in the array.

*
*/
func main() {
	fmt.Println(sumArr([]int{1, 2, 3, 4, 5}))
}

func sumArr(data []int) (total int) {
	for _, v := range data { //O(n)
		total = total + v
	}
	return
}
