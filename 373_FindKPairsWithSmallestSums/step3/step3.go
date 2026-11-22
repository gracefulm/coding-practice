package step3

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
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return nil
	}

	initialSize := min(len(nums1), k)
	candidates := make(minHeap, 0, initialSize)
	for i := range initialSize {
		heap.Push(&candidates, indexPairAndSum{i: i, j: 0, sum: nums1[i] + nums2[0]})
	}

	smallestPairs := make([][]int, 0, k)
	for candidates.Len() > 0 && len(smallestPairs) < k {
		p := heap.Pop(&candidates).(indexPairAndSum)
		smallestPairs = append(smallestPairs, []int{nums1[p.i], nums2[p.j]})
		if next := p.j + 1; next < len(nums2) {
			heap.Push(&candidates, indexPairAndSum{i: p.i, j: next, sum: nums1[p.i] + nums2[next]})
		}
	}
	return smallestPairs
}
