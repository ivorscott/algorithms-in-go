// find or calculate something among all the contiguous subarrays (or sublists) of a given size.
package main

import "fmt"

// Problem: Given an array find the average of all contiguous subarrays of size k in it.

// 1. Understand the problem with real input
//
// Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5

// Here, we are asked to find the average of all contiguous subarrays of size ‘5’ in the given array. Let’s solve this:

// For the first 5 numbers (subarray from index 0-4), the average is: (1+3+2+6-1)/5 => 2.2(1+3+2+6−1)/5=>2.2
// The average of next 5 numbers (subarray from index 1-5) is: (3+2+6-1+4)/5 => 2.8(3+2+6−1+4)/5=>2.8
// For the next 5 numbers (subarray from index 2-6), the average is: (2+6-1+4+1)/5 => 2.4(2+6−1+4+1)/5=>2.4

//Output: [2.2, 2.8, 2.4, 3.6, 2.8]

func main() {
	r := findAverageOfSubarrays(5, []int{1, 3, 2, 6, -1, 4, 1, 8, 2})
	fmt.Println(r)
}

// 2. Determine the brute force approach

func findAverageOfSubarrays(k int, arr []int) []float64 { // O (n * k) => n^2
	results := make([]float64, 0, k)
	// find sum of the next k elements
	for i := 0; i < len(arr)-k+1; i++ { // O(n) = 5
		var sum float64
		for j := i; j < i+k; j++ { // O(k) = 5
			sum += float64(arr[j])
		}
		results = append(results, sum/float64(k))
	}
	return results
}

// 3. Improve it

// The inefficiency is that for any two consecutive subarray of size 5, the
// overlapping part (which contains four elements) will be evaluated twice.

// Can we somehow reuse the sum we have calculated for the overlapping elements?

// visualize each contiguous subarray as a sliding window of 5 elements

// Slide the window by one element when we move to the next subarray
// To reuse the sum from the previous subarray, subtract the element going out of the window and
// add the next element to be included in the sliding window
func betterFindAverageOfSubarrays(k int, arr []int) []float64 { // O (n)
	var windowStart int
	var windowSum float64
	result := make([]float64, 0, k)

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ { // O(9)
		windowSum += float64(arr[windowEnd]) // add the next element
		// slide the window, we don't need to slide if we've not hit the required window size of k
		if windowEnd >= k-1 {
			result = append(result, windowSum/float64(k)) // calculate the average
			windowSum -= float64(arr[windowStart]) // subtract the element going out
			windowStart += 1 // slide the window ahead
		}
	}
	return result
}

