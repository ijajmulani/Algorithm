// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	arr := []int{1, 3, 1, 2, 10, 4}
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(first, second []int) []int {
	r := []int{}
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			r = append(r, first[i])
			i++
		} else {
			r = append(r, second[j])
			j++
		}
	}
	for i < len(first) {
		r = append(r, first[i])
		i++
	}
	for j < len(second) {
		r = append(r, second[j])
		j++
	}
	return r
}
