package step3

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(v any)        { *h = append(*h, v.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[h.Len()-1]
	*h = old[:h.Len()-1]
	return x
}

type KthLargest struct {
	k    int
	heap *minHeap
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k, heap: &minHeap{}}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	for this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}
	return (*this.heap)[0]
}
