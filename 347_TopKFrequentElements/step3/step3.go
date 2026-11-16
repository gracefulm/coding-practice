package step3

import "container/heap"

type numFrequency struct {
	num       int
	frequency int
}

type minHeap []numFrequency

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(numFrequency)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 {
		return nil
	}

	numToFrequency := make(map[int]int)
	for _, num := range nums {
		numToFrequency[num]++
	}

	topKHeap := make(minHeap, 0, k+1)
	for num, frequency := range numToFrequency {
		heap.Push(&topKHeap, numFrequency{num: num, frequency: frequency})
		if topKHeap.Len() > k {
			heap.Pop(&topKHeap)
		}
	}

	topK := make([]int, 0, topKHeap.Len())
	for topKHeap.Len() > 0 {
		topK = append(topK, heap.Pop(&topKHeap).(numFrequency).num)
	}

	return topK
}
