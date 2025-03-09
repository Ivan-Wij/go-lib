package utils

import (
	"cmp"
	stdHeap "container/heap"
)

type heap[T cmp.Ordered] []T

func (h heap[T]) Len() int           { return len(h) }
func (h heap[T]) Less(i, j int) bool { return h[i] < h[j] }
func (h heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Heap[T cmp.Ordered] struct {
	innerHeap *heap[T]
}

func (h *Heap[T]) Push(x T) {
	stdHeap.Push(h.innerHeap, x)
}

func (h *Heap[T]) Pop() T {
	return stdHeap.Pop(h.innerHeap).(T)
}

func (h *Heap[T]) IsEmpty() bool {
	return h.innerHeap.Len() == 0
}

func NewHeap[T cmp.Ordered](initialSize ...int) *Heap[T] {
	var innerHeap heap[T]
	if len(initialSize) == 1 {
		innerHeap = make([]T, initialSize[0])
	} else if len(initialSize) == 2 {
		innerHeap = make([]T, initialSize[0], initialSize[1])
	} else {
		innerHeap = make([]T, 0)
	}

	return &Heap[T]{innerHeap: &innerHeap}
}
