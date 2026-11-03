package step3

import "slices"

type KthLargest struct {
	k      int
	sorted []int
}

func Constructor(k int, nums []int) KthLargest {
	s := append([]int{}, nums...)
	slices.Sort(s)
	if len(s) > k {
		s = s[len(s)-k:]
	}
	return KthLargest{k: k, sorted: s}
}

func (this *KthLargest) Add(val int) int {
	i, _ := slices.BinarySearch(this.sorted, val)
	this.sorted = slices.Insert(this.sorted, i, val)
	if len(this.sorted) > this.k {
		this.sorted = this.sorted[1:]
	}
	return this.sorted[0]
}
