package main

import "fmt"

/*
	Problem:
	Find a value in a sorted array using the binary search algorithm.

*
*/
func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7}, 5))
}

func binarySearch(data []int, value int) (isFound bool) {
	for low, high := 0, len(data)-1; low <= high; {
		mid := (low + high) / 2
		if data[mid] == value {
			return true
		} else {
			if data[mid] > value {
				high = mid - 1 // move left
			} else {
				low = mid + 1 // move right
			}
		}
	}
	return
}
