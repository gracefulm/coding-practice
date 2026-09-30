package step2

import "slices"

type KthLargest struct {
	sortedScores []int
	k            int
}

func Constructor(k int, nums []int) KthLargest {
	ss := append([]int{}, nums...)
	slices.Sort(ss)
	if k < len(ss) {
		ss = ss[len(ss)-k:]
	}
	return KthLargest{sortedScores: ss, k: k}
}

func (this *KthLargest) Add(val int) int {
	i, _ := slices.BinarySearch(this.sortedScores, val)
	this.sortedScores = slices.Insert(this.sortedScores, i, val)
	if len(this.sortedScores) > this.k {
		this.sortedScores = this.sortedScores[1:]
	}
	return this.sortedScores[0]
}
