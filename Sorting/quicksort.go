// You can edit this code!
// Worst case o(n^2)
package main

import "fmt"

func main() {
	//	quicksort
	num := []int{10, 30, 20, 90, 5, 80}
	quickSort(num, 0, len(num)-1)
	fmt.Println(num)
}

func quickSort(num []int, low, high int) {
	if low < high {
		i := partitioning(num, low, high)
		quickSort(num, low, i-1)
		quickSort(num, i+1, high)
	}

}

func partitioning(num []int, low, high int) int {
	pivot := num[high]
	i := low
	for j := low; j < high; j++ {
		if num[j] < pivot {
			if i != j {
				num[i], num[j] = num[j], num[i]
			}
			i++
		}
	}
	num[i], num[high] = num[high], num[i]
	return i
}
