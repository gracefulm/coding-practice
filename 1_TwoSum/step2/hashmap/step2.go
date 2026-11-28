package hashmap

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numToIndex[target-num]; ok {
			return []int{j, i}
		}
		numToIndex[num] = i
	}
	return nil
}
