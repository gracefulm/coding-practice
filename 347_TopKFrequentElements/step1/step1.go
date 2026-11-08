package step1

import "slices"

type numFrequency struct {
	num       int
	frequency int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	numToFrequency := make(map[int]int, len(nums))
	for _, num := range nums {
		f := 1
		if c, ok := numToFrequency[num]; ok {
			f = c + 1
		}
		numToFrequency[num] = f
	}

	numFrequencies := make([]numFrequency, 0, k)
	for num, frequency := range numToFrequency {
		numFrequencies = append(numFrequencies, numFrequency{num: num, frequency: frequency})
	}
	slices.SortFunc(numFrequencies, func(i, j numFrequency) int {
		return j.frequency - i.frequency
	})

	topK := make([]int, 0, k)
	if len(nums) > k {
		numFrequencies = numFrequencies[:k]
	}
	for _, nf := range numFrequencies {
		topK = append(topK, nf.num)
	}
	return topK
}
