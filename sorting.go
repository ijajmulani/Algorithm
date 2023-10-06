package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		min := i
		for j := i; j < n; j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]	
				swap = true
			}
		}
		fmt.Println(i)
		if !swap {
			break
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 6, 5, 4, 7, 8, 9}
	bubbleSort(arr)
	fmt.Println(arr)
}
