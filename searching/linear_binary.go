// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

// O(n) worst case time complexity
// O(1)  space complexity
// SLow for large dataset
func LinearSearch(arr []int, s int) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}

// O(log n) worst case time complexity
// O(1)  space complexity
// large dataset 

func BinarySearch(arr []int, s int, l, r int) int {
	sort.Ints(arr)

	for l <= r {

		mid := l + (r-l)/2
		if arr[mid] == s {
			return mid
		} else if s > arr[mid] {
			l = mid + 1
		} else if s < arr[mid] {
			r = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{7, 2, 3, 4, 5}
	fmt.Println(LinearSearch(arr, 4))
	fmt.Println(BinarySearch(arr, 7, 0, len(arr)-1))
}
