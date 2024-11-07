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
func (h *maxHeap) Heapify(i int, n int) {
	largest := i
	l := h.Left(i)
	r := h.Right(i)
	if l < n && h.data[l] > h.data[largest] {
		largest = l
	}
	if r < n && h.data[r] > h.data[largest] {
		largest = r
	}
	if largest != i {
		h.data[largest], h.data[i] = h.data[i], h.data[largest]
		h.Heapify(largest, n)
	}
}

// o(n)
func (h *maxHeap) BuildHeap(arr []int) {
	h.data = arr
	for i := h.Parent(len(h.data) - 1); i >= 0; i-- {
		h.Heapify(i, len(h.data))
	}
}

func (h *maxHeap) Sort() {
	for i := len(h.data) - 1; i > 0; i-- {
		h.data[i], h.data[0] = h.data[0], h.data[i]
		h.Heapify(0, i)
	}
}

func (h *maxHeap) Print() {
	fmt.Println(h.data)
}

func main() {
	h := new(maxHeap)
	buildHeap := []int{10, 15, 50, 4, 20}
	h.BuildHeap(buildHeap)
	
	h.Sort()
	h.Print()
}
