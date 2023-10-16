


type maxHeap struct {
    data []heapData
    size int
}
type heapData struct{
        key int
        val int
    }
    
func(h *maxHeap) Parent(i int) int {
    return (i-1)/2
}
func(h *maxHeap) Left(i int)int {
    return 2*i +1
}
func(h *maxHeap) Right(i int) int{
    return 2*i +2
}
func(h *maxHeap) Heapify(i int) {
    smallest := i
    l:=h.Left(i)
    r:=h.Right(i)
    n := len(h.data)
    if l < n && h.data[l].key > h.data[smallest].key  {
        smallest = l
    }
    if r < n && h.data[r].key > h.data[smallest].key  {
        smallest = r
    }
    if smallest !=i {
        h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
        h.Heapify(smallest)
    }
}

func(h *maxHeap) Pop() int {
    root := h.data[0].val
    h.data[0], h.data[len(h.data) -1] = h.data[len(h.data) -1], h.data[0]
    h.data=h.data[:len(h.data) -1]
    h.Heapify(0)
    return root
}

func(h *maxHeap) Push(d heapData) {
    if len(h.data) == 0 {
        h.data = append(h.data, d)
        return
    }
    if len(h.data) == h.size  && (d.key == h.data[0].key && d.val > h.data[0].val) {
        return
    } else if len(h.data) == h.size  && d.key > h.data[0].key {
        return 
    } else if len(h.data) == h.size {
        h.Pop()
    }
    
    h.data = append(h.data, d)
    i:= len(h.data) -1
    for i>=0 && h.data[i].key > h.data[h.Parent(i)].key   {
        h.data[i], h.data[h.Parent(i)] = h.data[h.Parent(i)], h.data[i]
        i=h.Parent(i)
    }
}
func(h *maxHeap) Print() {
for _,d:= range h.data{
    fmt.Printf("(k = %d v= %d),  ", d.key, d.val)
}
fmt.Println("")
}


func findClosestElements(arr []int, k int, x int) []int {
    h := new(maxHeap)
    h.size = k
  
    for _,v:= range arr{
        diff := v-x
        if diff < 0 {
            diff = -diff
        }
        
        data := heapData{
            key:  diff,
            val: v,
        }
        h.Push(data)
    }
    var result  []int
    for len(h.data) > 0 {
        result = append(result, h.Pop())
    } 
    
    sort.Slice(result, func(i int, j int) bool {
      return result[i] < result[j]
    })
    return result
}
