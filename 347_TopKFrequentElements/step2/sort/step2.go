package step2

import (
	"cmp"
	"slices"
)

type numFrequency struct {
	num       int
	frequency int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return nil
	}

	frequency := make(map[int]int, len(nums))
	for _, n := range nums {
		frequency[n]++
	}

	nfs := make([]numFrequency, 0, len(frequency))
	for n, f := range frequency {
		nfs = append(nfs, numFrequency{num: n, frequency: f})
	}
	slices.SortFunc(nfs, func(i, j numFrequency) int {
		return cmp.Compare(j.frequency, i.frequency)
	})

	topK := make([]int, 0, k)
	nfs = nfs[:min(len(nfs), k)]
	for _, nf := range nfs {
		topK = append(topK, nf.num)
	}
	return topK
}
