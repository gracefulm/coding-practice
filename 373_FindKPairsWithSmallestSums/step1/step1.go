package step1

import "container/heap"

type twoNums struct {
	num1 int
	num2 int
}

func (tn twoNums) sum() int {
	return tn.num1 + tn.num2
}

func (tn twoNums) toSlice() []int {
	return []int{tn.num1, tn.num2}
}

type minHeap []twoNums

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i].sum() < mh[j].sum() }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *minHeap) Push(x any)        { *mh = append(*mh, x.(twoNums)) }
func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	mh := make(minHeap, 0, k)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			heap.Push(&mh, twoNums{num1: num1, num2: num2})
		}
	}
	l := min(k, mh.Len())
	smallestPairs := make([][]int, 0, l)
	for range l {
		smallestPairs = append(smallestPairs, heap.Pop(&mh).(twoNums).toSlice())
	}
	return smallestPairs
}
