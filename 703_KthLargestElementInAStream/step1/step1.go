package step1

import "sort"

type KthLargest struct {
	k int
	// sortedScores は、降順ソートされたスコアです。
	sortedScores []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Slice(nums, func(i int, j int) bool { return nums[i] > nums[j] })
	return KthLargest{k: k, sortedScores: nums}
}

func (this *KthLargest) Add(val int) int {
	// すでにソートされたスライスに要素を追加する特性を考慮してより良いソートアルゴリズムを選択できる可能性はある
	this.sortedScores = append(this.sortedScores, val)
	sort.Slice(this.sortedScores, func(i, j int) bool { return this.sortedScores[i] > this.sortedScores[j] })

	// 要件次第だが、とりあえず要素数がkに満たない場合は最大値を返す。
	if len(this.sortedScores) < this.k {
		return this.sortedScores[0]
	}

	return this.sortedScores[this.k-1]
}
