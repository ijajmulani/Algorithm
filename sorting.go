package main

import (
	"fmt"
	"math/rand"
	"time"
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
		if !swap {
			break
		}
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func mergeSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}
	mid := n / 2

	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])

	i, j := 0, 0
	final := []int{}
	for i < len(l) && j < len(l) {
		if l[i] < r[j] {
			final = append(final, l[i])
			i++
		} else {
			final = append(final, r[j])
			j++
		}
	}

	for ; i < len(l); i++ {
		final = append(final, l[i])
	}
	for ; j < len(r); j++ {
		final = append(final, r[j])
	}
	return final
}

func main() {
	rand.Seed(time.Now().Unix())
	
	arr := rand.Perm(10000)
	start := time.Now()
	selectionSort(arr)
	fmt.Printf(" selectionSort took at time %v\n", time.Since(start))
	
	arr=nil
	arr = rand.Perm(10000)
	start = time.Now()
	bubbleSort(arr)
	fmt.Printf(" bubbleSort took at time %v\n", time.Since(start))
	
	arr=nil
	arr = rand.Perm(10000)
	start = time.Now()
	insertionSort(arr)
	fmt.Printf(" insertionSort took at time %v\n", time.Since(start))
	
	arr=nil
	arr = rand.Perm(10000)
	start = time.Now()
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf(" quickSort took at time %v\n", time.Since(start))
}

func partition(arr []int, low, high int) int {
	key := arr[high]
	j := low
	for i := low; i < high; i++ {
		if arr[i] < key {

			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	arr[j], arr[high] = arr[high], arr[j]
	return j
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)
}
