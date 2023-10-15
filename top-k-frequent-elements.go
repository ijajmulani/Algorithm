// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	a := topKFrequent([]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}, 3)
	fmt.Println(a)
}

type kvstruct struct {
	key       int
	occurance int
}
type minHeap struct {
	data []kvstruct
	size int
}

func (h *minHeap) Heapify(i int) {
	smallest := i
	l := (2 * i) + 1
	r := (2 * i) + 2
	n := len(h.data)
	if l < n && h.data[l].occurance < h.data[smallest].occurance {
		smallest = l
	}
	if r < n && h.data[r].occurance < h.data[smallest].occurance {
		smallest = r
	}
	if i != smallest {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.Heapify(smallest)
	}
}

func (h *minHeap) ExtractRoot() {
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.Heapify(0)
}
func (h *minHeap) Print() {
	for _, v := range h.data {
		fmt.Printf("%d ", v.occurance)

	}
	fmt.Println("")
}

func (h *minHeap) Insert(k, v int) {
	if len(h.data) == 0 {
		h.data = append(h.data, kvstruct{key: k, occurance: v})
		return
	}
	if len(h.data) == h.size && v < h.data[0].occurance {
		return
	}

	if len(h.data) == h.size {
		h.ExtractRoot()
	}
	h.data = append(h.data, kvstruct{key: k, occurance: v})
	i := len(h.data) - 1
	for i >= 0 && h.data[i].occurance < h.data[(i-1)/2].occurance {
		h.data[i], h.data[(i-1)/2] = h.data[(i-1)/2], h.data[i]
		i = (i - 1) / 2
	}
}

func topKFrequent(nums []int, k int) []int {
	unorderedMap := make(map[int]int)
	for _, v := range nums {
		unorderedMap[v] = unorderedMap[v] + 1
	}

	h := new(minHeap)
	h.size = k
	for k, v := range unorderedMap {
		h.Insert(k, v)
	}
	var r []int
	for _, v := range h.data {
		r = append(r, v.key)
	}
	return r
}
