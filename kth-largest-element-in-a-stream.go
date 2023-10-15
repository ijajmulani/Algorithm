// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	a := Constructor(3, []int{4, 5, 8, 2})
	var r int
	r = a.Add(3)
	fmt.Println(r)
	r = a.Add(5)
	fmt.Println(r)
	r = a.Add(10)
	fmt.Println(r)
	r = a.Add(9)
	fmt.Println(r)
	r = a.Add(4)
	fmt.Println(r)
}

type KthLargest struct {
	minheap []int
	size    int
}

func Constructor(k int, nums []int) KthLargest {
	h := new(KthLargest)
	h.size = k
	for _, v := range nums {
		h.Add(v)
	}

	return *h
}

func (this *KthLargest) Heapify(val int) {
	smallest := val
	l := 2*val + 1
	r := 2*val + 2
	n := len(this.minheap)
	if l < n && this.minheap[l] < this.minheap[smallest] {
		smallest = l
	}
	if r < n && this.minheap[r] < this.minheap[smallest] {
		smallest = r
	}

	if val != smallest {
		this.minheap[smallest], this.minheap[val] = this.minheap[val], this.minheap[smallest]
		this.Heapify(smallest)
	}
}

func (this *KthLargest) ExtractRoot() {
	this.minheap[0], this.minheap[len(this.minheap)-1] = this.minheap[len(this.minheap)-1], this.minheap[0]
	this.minheap = this.minheap[:len(this.minheap)-1]
	this.Heapify(0)
}

func (this *KthLargest) Add(val int) int {
	if len(this.minheap) == 0 {
		this.minheap = append(this.minheap, val)
		return this.minheap[0]
	}
	if len(this.minheap) == this.size && val < this.minheap[0] {
		return this.minheap[0]
	}
	if len(this.minheap) == this.size {
		this.ExtractRoot()
	}

	this.minheap = append(this.minheap, val)
	i := len(this.minheap) - 1
	for i > 0 && this.minheap[i] <= this.minheap[(i-1)/2] {
		this.minheap[(i-1)/2], this.minheap[i] = this.minheap[i], this.minheap[(i-1)/2]
		i = (i - 1) / 2
	}
	return this.minheap[0]
}
