// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

type heap struct {
	maxSize int
	data    []int
}

func (h *heap) parent(k int) int {
	return int(math.Floor(float64((k - 1) / 2)))
}

func (h *heap) lChildIndex(k int) int {
	return 2*k + 1
}
func (h *heap) rChildIndex(k int) int {
	return 2*k + 2
}

func (h *heap) insert(v int) {
	if len(h.data) == h.maxSize {
		return
	}
	h.data = append(h.data, v)
	i := len(h.data) - 1
	for i != 0 && h.data[h.parent(i)] < h.data[i] {
		h.data[h.parent(i)], h.data[i] = h.data[i], h.data[h.parent(i)]
		i = h.parent(i)
	}
}

func (h *heap) increaseKey(i int) {
	h.data[i] = math.MaxInt32
	for i != 0 && h.data[h.parent(i)] < h.data[i] {
		h.data[h.parent(i)], h.data[i] = h.data[i], h.data[h.parent(i)]
		i = h.parent(i)
	}
}

func (h *heap) heapify(i int) {

	lc := h.lChildIndex(i)
	rc := h.rChildIndex(i)
	largest := i
	if lc < len(h.data) && h.data[lc] > h.data[largest] {
		largest = lc
	}
	if rc < len(h.data) && h.data[rc] > h.data[largest] {
		largest = rc
	}
	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
	}

}

func (h *heap) removeMax() {
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapify(0)

}

func (h *heap) delete(i int) {
	if len(h.data) == 0 || (i == 1 && len(h.data) == 1) {
		h.data = nil
		return
	}

	if i <= len(h.data)-1 {
		h.increaseKey(i)
		h.removeMax()
	}
}

func main() {
	h := &heap{
		maxSize: 15,
	}
	h.insert(4)
	h.insert(1)
	h.insert(3)
	h.insert(7)
	fmt.Println(h.data, len(h.data))
	h.delete(2)
	fmt.Println(h.data)
}
