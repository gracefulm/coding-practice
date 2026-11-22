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
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return nil
	}

	// nums1, nums2 はソート済みなので、(i, j) より先に (i, j+1) が答えになることはない。
	// そのため各 i について (i, 0) から始め、(i, j) を取り出したときに (i, j+1) を候補に加える。
	initialSize := min(len(nums1), k)
	candidates := make(minHeap, 0, initialSize)
	for i := range initialSize {
		heap.Push(&candidates, indexPairAndSum{i: i, j: 0, sum: nums1[i] + nums2[0]})
	}

	smallestPairs := make([][]int, 0, k)
	for len(smallestPairs) < k && candidates.Len() > 0 {
		smallest := heap.Pop(&candidates).(indexPairAndSum)
		smallestPairs = append(smallestPairs, []int{nums1[smallest.i], nums2[smallest.j]})
		if next := smallest.j + 1; next < len(nums2) {
			heap.Push(&candidates, indexPairAndSum{i: smallest.i, j: next, sum: nums1[smallest.i] + nums2[next]})
		}
	}
	return smallestPairs
}
