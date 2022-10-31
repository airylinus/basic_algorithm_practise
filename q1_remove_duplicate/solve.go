package q1_remove_duplicate

func removeDuplicate(nums []int) int {
	re, pn := 0, 1
	for pn < len(nums) {
		// if pn equal to the end value of result, move pn to next
		if nums[re] == nums[pn] {
			pn += 1
			continue
		}
		// pn != re
		// add pn to result end
		re += 1
		nums[re] = nums[pn]
		pn += 1
	}
	return re + 1
}
