// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type maxHeap struct {
	data []int
}

func (h *maxHeap) Left(i int) int {
	return 2*i + 1
}
func (h *maxHeap) Right(i int) int {
	return 2*i + 2
}
func (h *maxHeap) Parent(i int) int {
	return (i - 1) / 2
}

// o(log n)
func (h *maxHeap) Heapify(i int) {
	largest := i
	l := h.Left(i)
	r := h.Right(i)
	n := len(h.data)
	if l < n && h.data[l] > h.data[largest] {
		largest = l
	}
	if r < n && h.data[r] > h.data[largest] {
		largest = r
	}
	if largest != i {
		h.data[largest], h.data[i] = h.data[i], h.data[largest]
		h.Heapify(largest)
	}
}

// o(n)
func (h *maxHeap) BuildHeap(arr []int) {
	h.data = arr
	for i := h.Parent(len(h.data) - 1); i >= 0; i-- {
		h.Heapify(i)
	}
}
func (h *maxHeap) ExtractRoot() int {
	root := h.data[0]
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[:len(h.data)-1]
	h.Heapify(0)
	return root
}

func (h *maxHeap) Print() {
	fmt.Println(h.data)
}

func main() {
	h := new(maxHeap)
	buildHeap := []int{3, 2, 1, 5, 6, 4}
	h.BuildHeap(buildHeap)
	fmt.Println(h.data)
	k := 2
	for i := 0; i < k-1; i++ {
		h.ExtractRoot()
	}
	fmt.Println(h.data[0])
}
