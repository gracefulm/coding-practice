package step1

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	for i, num1 := range nums[:len(nums)-1] {
		for j, num2 := range nums[i+1:] {
			if target == num1+num2 {
				return []int{i, i + 1 + j}
			}
		}
	}
	return nil
}
