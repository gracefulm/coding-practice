package heap

import "container/heap"

type numFreq struct {
	num  int
	freq int
}

type minHeap []numFreq

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(numFreq))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// heap を使った実装
//
// 時間計算量: O(n + m log k)
//   - map の構築: O(n)
//   - heap への追加・削除: O(m log k)
//   - m: unique elements の数
//   - k: 上位 k 個
//
// 空間計算量: O(m)
//   - map: O(m)
//   - heap: O(k)
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 {
		return nil
	}

	numsFreq := make(map[int]int)
	for _, num := range nums {
		numsFreq[num]++
	}

	mh := make(minHeap, 0, k+1)
	for num, freq := range numsFreq {
		heap.Push(&mh, numFreq{num: num, freq: freq})
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
