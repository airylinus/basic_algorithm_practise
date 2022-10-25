package q8_move_zero

// moveZero 移动零
// 1,0,2,3
func moveZero(nums []int) []int {
	var n, z int
	// 先从前向后找到 0 的数字
	for z < len(nums) {
		// 不是 0 的跳过
		if nums[z] != 0 {
			z += 1
			continue
		}
		// 找到 0 ，从 0 开始向后找非零数字
		n = z
		for n < len(nums) {
			// 也是 0 ，跳过，继续找
			if nums[n] == 0 {
				n += 1
				continue
			}
			// 找到非 0 数字，跟前面的 0 交换位置
			nums[n], nums[z] = nums[z], nums[n]
			break
		}
		// 交换之后，从下一个位置开始继续找 0 的数字
		z += 1
	}
	return nums
}
