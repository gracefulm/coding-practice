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

	frequency := make(map[int]int, len(nums))
	for _, n := range nums {
		f := 1
		if c, ok := frequency[n]; ok {
			f = c + 1
		}
		frequency[n] = f
	}

	sortedNFs := make([]numFrequency, 0, k)
	for n, f := range frequency {
		sortedNFs = append(sortedNFs, numFrequency{num: n, frequency: f})
	}
	slices.SortFunc(sortedNFs, func(i, j numFrequency) int {
		return j.frequency - i.frequency
	})

	topK := make([]int, 0, k)
	if len(nums) > k {
		sortedNFs = sortedNFs[:k]
	}
	for _, nf := range sortedNFs {
		topK = append(topK, nf.num)
	}
	return topK
}
