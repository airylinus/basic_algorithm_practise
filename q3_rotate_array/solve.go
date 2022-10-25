package q3_rotate_array

import "fmt"

/*
 * 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
 *输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
*/

// rotateArray
// 最后的 K 个元素会跑到最前面来，
// 1. 整个翻转数组可以把后 K 个数字调到头部。  [7,6,5,4,3,2,1]
// 2. 经过 1 之后前面 K 个翻转一次，即可得到前面部分的正确结果. [5,6,7,4,3,2,1]
// 3. 再翻转后面的 len - k 个元素即可得到答案   [5,6,7,1,2,3,4]
func rotateArraySolve1(nums []int, k int) (r []int) {
	step := k % len(nums)
	revertNums(nums)
	fmt.Println(nums)
	revertNums(nums[:step])
	fmt.Println(nums)
	revertNums(nums[step:])
	//revertNums(nums[:])
	return nums
}

func revertNums(nums []int) []int {
	h := 0
	e := len(nums) - 1
	for h < e {
		nums[h], nums[e] = nums[e], nums[h]
		h += 1
		e -= 1
	}
	return nums
}
