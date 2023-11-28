package main

import "fmt"

/*
	Problem:
	Write a method that searches for a given value in an unsorted array.

*
*/
func main() {
	fmt.Println(sequentialSearch([]int{10, 1, 2, 3, 4, 5}, 5))
}

func sequentialSearch(data []int, value int) (isFound bool) {
	for _, v := range data { //O(n)
		if value == v {
			return true
		}
	}
	return //Return false if value not found
}
