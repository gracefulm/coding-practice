package step2

import "container/heap"

type indexPairAndSum struct {
	i, j int
	sum  int
}

type minHeap []indexPairAndSum

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i].sum < mh[j].sum }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *minHeap) Push(x any)        { *mh = append(*mh, x.(indexPairAndSum)) }
func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	mh := make(minHeap, 0, min(len(nums1), k))
	for i := range min(len(nums1), k) {
		heap.Push(&mh, indexPairAndSum{i: i, j: 0, sum: nums1[i] + nums2[0]})
	}
	smallestPairs := make([][]int, 0, k)
	for len(smallestPairs) < k && mh.Len() > 0 {
		p := heap.Pop(&mh).(indexPairAndSum)
		smallestPairs = append(smallestPairs, []int{nums1[p.i], nums2[p.j]})
		if p.j+1 < len(nums2) {
			heap.Push(&mh, indexPairAndSum{i: p.i, j: p.j + 1, sum: nums1[p.i] + nums2[p.j+1]})
		}
	}
	return smallestPairs
}
