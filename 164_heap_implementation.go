// implementing maxheap using container/heap package
// code hand-copied from chatGPT , but changed variable names etc and then troubleshooting coding mistakes
// JAN 22, 2025

package main

import (
"container/heap"
"fmt"
)

type maxheap []int

// WE MUST IMPLEMENT 5 func for us to use the heap interface
// it is requirement of THAT INTERFACE
// Len
// Less
// Swap
// Push
// Pop

func (h maxheap) Len() int {
return len(h)
}

func (h maxheap) Less(i, j int) bool {
if h[i] > h[j] { // logic inverted because it is Max heap, default is a Min heap
return true
} else {
return false
}
}

func (h maxheap) Swap(i, j int) {
h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Push(x any) {
*h = append(*h, x.(int))
}

func (h *maxheap) Pop() any {
old := *h
n := len(old)
x := old[n-1]
*h = old[0 : n-1]
return x

}

func main() {

heap1 := &maxheap{100, 900, 300, 500, 400, 1, 10000, 20, 5000, 70}
fmt.Println(*heap1) // [100 900 300 500 400 1 10000 20 5000 70]
heap.Init(heap1)    // like heapify in Python
fmt.Println(*heap1) // [10000 5000 300 900 400 1 100 20 500 70] largest at the top
heap.Pop(heap1)
fmt.Println(*heap1) // [5000 900 300 500 400 1 100 20 70] , because this is maxheap, largest will be popped
heap.Push(heap1, 2000)
fmt.Println(*heap1) // [5000 2000 300 500 900 1 100 20 70 400] , the item "2000" got added at the right spot

}

