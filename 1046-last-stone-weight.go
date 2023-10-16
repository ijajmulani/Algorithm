// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	lastStoneWeight([]int{2, 7, 4, 1, 8, 1})

}

type maxHeap []int

func (h *maxHeap) Parent(i int) int {
	return (i - 1) / 2
}
func (h *maxHeap) Left(i int) int {
	return 2*i + 1
}
func (h *maxHeap) Right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) Heapify(i int) {
	largest := i
	l := h.Left(i)
	r := h.Right(i)
	n := len(*h)
	if l < n && (*h)[l] > (*h)[largest] {
		largest = l
	}
	if r < n && (*h)[r] > (*h)[largest] {
		largest = r
	}
	if largest != i {
		(*h)[largest], (*h)[i] = (*h)[i], (*h)[largest]
		i = largest
		h.Heapify(i)
	}
}

func (h *maxHeap) Pop() int {
	root := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	(*h) = (*h)[:len(*h)-1]
	h.Heapify(0)
	return root
}
func (h *maxHeap) Push(v int) {
	*h = append(*h, v)
	i := len(*h) - 1

	for i > 0 && (*h)[h.Parent(i)] < (*h)[i] {
		(*h)[h.Parent(i)], (*h)[i] = (*h)[i], (*h)[h.Parent(i)]
		i = h.Parent(i)
	}
}

func lastStoneWeight(stones []int) int {
	h := new(maxHeap)
	for _, s := range stones {
		h.Push(s)
	}
	for len(*h) > 1 {
		f := h.Pop()
		s := h.Pop()
        if f != s {
            h.Push(f - s)
        }
	}
    if len((*h)) > 0 {
        return (*h)[0]
    }
	return 0
}
