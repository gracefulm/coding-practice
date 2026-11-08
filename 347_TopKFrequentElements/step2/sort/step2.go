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

	numToFrequency := make(map[int]int, len(nums))
	for _, num := range nums {
		numToFrequency[num]++
	}

	numFrequencies := make([]numFrequency, 0, len(numToFrequency))
	for num, frequency := range numToFrequency {
		numFrequencies = append(numFrequencies, numFrequency{num: num, frequency: frequency})
	}
	slices.SortFunc(numFrequencies, func(i, j numFrequency) int {
		return cmp.Compare(j.frequency, i.frequency)
	})

	topK := make([]int, 0, k)
	numFrequencies = numFrequencies[:min(len(numFrequencies), k)]
	for _, nf := range numFrequencies {
		topK = append(topK, nf.num)
	}
	return topK
}
