// implementing maxheap using container/heap package
// code hand-copied from chatGPT , but changed variable names etc and then troubleshooting coding mistakes
// JAN 22, 2025

package main

import (
"container/heap"
"fmt"
)


// to have custom methods, we must define a new TYPE from Slice
type maxheap []int

// WE MUST IMPLEMENT 5 func for us to use the heap interface
// it is requirement of THAT INTERFACE
// Len
// Less
// Swap
// Push
// Pop


// LEN, easy , takes in he "TYPE" = maxheap
func (h maxheap) Len() int {
    return len(h)
}

//LESS : is i less or j (indexes)  less?? that's the question that gets answered
func (h maxheap) Less(i, j int) bool {
  if h[i] > h[j] { // logic inverted because it is Max heap, default is a Min heap
    return true
  } else {
     return false
}
}

// SWAP, easy stuff!
func (h maxheap) Swap(i, j int) {
     h[i], h[j] = h[j], h[i]
}

// PUSH , EASY BUT takes is the DERERENCING of pointer to max heap == *maxheap
func (h *maxheap) Push(x any) {
   *h = append(*h, x.(int))
}

// POP ' This is the hardest one to remember
// takes in a pointer to a maxheap
// the idea is to point to the SECOND element (i.e. POP the first one
// to do that , we need do : *h = *h[0 : n-1]
//  BUT that will "lose" the top element
func (h *maxheap) Pop() any {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]

  return x

}

func main() {

// create a new heap and fill in data in one shot
// note the & sign
heap1 := &maxheap{100, 900, 300, 500, 400, 1, 10000, 20, 5000, 70}
fmt.Println(*heap1) // [100 900 300 500 400 1 10000 20 5000 70]

// heapify
heap.Init(heap1)    // like heapify in Python
fmt.Println(*heap1) // [10000 5000 300 900 400 1 100 20 500 70] largest at the top

// pop 1
heap.Pop(heap1)
fmt.Println(*heap1) // [5000 900 300 500 400 1 100 20 70] , because this is maxheap, largest will be popped
// push and gets added where it belongs
heap.Push(heap1, 2000)
fmt.Println(*heap1) // [5000 2000 300 500 900 1 100 20 70 400] , the item "2000" got added at the right spot

}

