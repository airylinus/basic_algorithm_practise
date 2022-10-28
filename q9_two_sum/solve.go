package q9_two_sum

func towSum(nums []int, target int) []int {
	r := make([]int, 2)
	valMap := make(map[int]int, 0)
	for k, v := range nums {
		if j, ok := valMap[target-v]; ok {
			r[0] = j
			r[1] = k
		}
		valMap[nums[k]] = k
	}
	return r
}
