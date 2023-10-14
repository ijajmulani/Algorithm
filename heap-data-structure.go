// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

type minHeap struct {
	data []int
}

func (h *minHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *minHeap) LeftNode(i int) int {
	return 2*i + 1
}

func (h *minHeap) RightNode(i int) int {
	return 2*i + 2
}

func (h *minHeap) Insert(v int) {
	h.data = append(h.data, v)
	n := len(h.data) - 1
	for n > 0 && h.data[h.Parent(n)] > h.data[n] {
		h.data[h.Parent(n)], h.data[n] = h.data[n], h.data[h.Parent(n)]
		n = h.Parent(n)
	}
}

func (h *minHeap) MinHeapify(i int) {
	smallest := i
	left := h.LeftNode(i)
	right := h.RightNode(i)
	if left < len(h.data) && h.data[left] < h.data[smallest] {
		smallest = left
	}
	if right < len(h.data) && h.data[right] < h.data[smallest] {
		smallest = right
	}
	if i != smallest {
		h.data[smallest], h.data[i] = h.data[i], h.data[smallest]
		h.MinHeapify(smallest)
	}
}

func (h *minHeap) ExtractMin() {
	root := 0
	last := len(h.data) - 1
	h.data[root], h.data[last] = h.data[last], h.data[root]
	h.data = h.data[:len(h.data)-1]
	h.MinHeapify(0)
}

func (h *minHeap) DeleteKey(i int) {
	h.data[i] = math.MinInt
	for i >= 0 && h.data[i] < h.data[h.Parent(i)] {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
	h.ExtractMin()
}

func (h *minHeap) BuildHeap() {
	i := h.Parent(len(h.data) - 1)
	for ; i >= 0; i-- {
		h.MinHeapify(i)
	}
}
func (h *minHeap) Print() {
	fmt.Println(h.data)
}

func main() {
	h := new(minHeap)
	// [10 20 15 40 50 100 25 45]

	h.Insert(40)
	h.Insert(20)
	h.Insert(30)
	h.Insert(35)
	h.Insert(25)
	h.Insert(80)
	h.Insert(32)
	h.Insert(100)
	h.Insert(70)
	h.Insert(60)

	//h.ExtractMin()
	h.Print()
	//h.DeleteKey(3)
	h.BuildHeap()
	h.Print()
}
