package step3

import "container/heap"

type numFreq struct {
	num  int
	freq int
}

type minHeap []numFreq

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i].freq < mh[j].freq }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *minHeap) Push(x any)        { *mh = append(*mh, x.(numFreq)) }
func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 {
		return nil
	}

	numFreqMap := make(map[int]int)
	for _, num := range nums {
		numFreqMap[num]++
	}

	mh := make(minHeap, 0, k+1)
	for n, f := range numFreqMap {
		heap.Push(&mh, numFreq{num: n, freq: f})
		if mh.Len() > k {
			heap.Pop(&mh)
		}
	}

	topk := make([]int, 0, mh.Len())
	for mh.Len() > 0 {
		topk = append(topk, heap.Pop(&mh).(numFreq).num)
	}

	return topk
}
