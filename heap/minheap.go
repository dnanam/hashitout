package main

import (
	"container/heap"
	"fmt"
)

/*

* GoLang provides an interface which can be implemented to produce the Heap DS
* under the hood; heaps are balanced binary tree

 */

// myIntHeap is of type int
// I need a type to implement the interface golang provides out of the box; basically I need to provide an implementation for the functions
// that heap will perform
type myIntHeap []int

func (h myIntHeap) Len() int {
	return len(h)
}

func (h myIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h myIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// all Push does is append to the end of the heap
func (h *myIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop removes the first element from the heap
func (h *myIntHeap) Pop() any {
	oldHeap := *h
	size := len(oldHeap)
	toPop := oldHeap[size-1]
	*h = oldHeap[:size-1]
	return toPop
}

func main() {
	h := &myIntHeap{1, 2, 3}
	heap.Init(h)
	element := heap.Pop(h)
	fmt.Println("element pop'd", element)
	fmt.Println(h)
	heap.Push(h, 6)
	fmt.Println(h)

}
