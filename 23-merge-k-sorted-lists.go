/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type minHeap struct{
     data []*ListNode
     size int
 }

 func (h *minHeap) Parent(i int) int{
     return (i-1)/2
 }
 func (h *minHeap) Left(i int) int{
     return 2*i + 1
 }
 func (h *minHeap) Right(i int) int{
     return 2*i + 2
 }

 func (h *minHeap) Push(l *ListNode) {
    if len(h.data) == 0 {
        h.data = append(h.data, l)
        return
    }

    h.data = append(h.data, l)
    i := len(h.data) -1
    for i > 0 && h.data[i].Val < h.data[h.Parent(i)].Val {
        h.data[i], h.data[h.Parent(i)] = h.data[h.Parent(i)], h.data[i]
        i = h.Parent(i)
    }   
 }

 func (h *minHeap) Pop() *ListNode {
     root := h.data[0]
     
     h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
     h.data = h.data[:len(h.data)-1]
     h.Heapify(0)
     return root
 }

func(h *minHeap) Heapify(i int) {
    smallest := i
    l:=h.Left(i)
    r:=h.Right(i)
    n := len(h.data)
    if l < n && h.data[l].Val < h.data[smallest].Val  {
        smallest = l
    }
    if r < n && h.data[r].Val < h.data[smallest].Val  {
        smallest = r
    }
    if smallest !=i {
        h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
        h.Heapify(smallest)
    }
}
func (h *minHeap) Print() {
    fmt.Println("Printing")
    fmt.Println(h.data)
    for _,l := range h.data {
        fmt.Printf(" => %d ", l.Val)
    }
    fmt.Println("")
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := new(minHeap)
    h.size = len(h.data)
    for _, l := range lists {
        if l != nil {
            h.Push(l)
        }
    }
    var ret *ListNode
    var prev *ListNode
    for len(h.data) > 0 {
        d := h.Pop()
        
        if prev == nil {
            prev = d
            ret = d
        } else {
            prev.Next = d
            prev = d
        }
       
        if d.Next != nil {
            h.Push(d.Next)
        }
    }
    return ret
}
