package q7_add_one

// addOne 给每一位数字加 1，错误版本
func addOne(nums []int) (rn []int) {
	odds := 0
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i] + odds + 1
		if n > 10 {
			odds = 1
			nums[i] = n % 10
		} else {
			odds = 0
			nums[i] = n
		}
	}
	rn = nums
	return
}

func plusOne(digits []int) (r []int) {
	odds := 0
	digits = revertArray(digits)
	//fmt.Println(digits)
	for i := 0; i < len(digits); i++ {
		n := digits[i] + odds
		if i == 0 {
			n += 1
		}
		if n >= 10 {
			digits[i] = n % 10
			odds = 1
		} else {
			digits[i] = n
			odds = 0
		}
	}
	if odds == 1 {
		digits = append(digits, odds)
	}
	return revertArray(digits)
}

func revertArray(nums []int) (r []int) {
	h := 0
	e := len(nums) - 1
	for h < e {
		nums[h], nums[e] = nums[e], nums[h]
		h += 1
		e -= 1
	}
	return nums
}
