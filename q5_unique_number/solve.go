package q5_unique_number

// getUniqueNumber 其他数字都是出现两次，只有一个数字只出现一次，求这个数字
// 例如 ： [ 1, 2, 2, 3, 3, 4, 4] 返回 1
//        [ 3, 3, 4, 4, 5] 返回 5
func getUniqueNumberVersion1(numbers []int) (num int) {
	uniqueMap := make(map[int]struct{})
	for _, n := range numbers {
		if _, ok := uniqueMap[n]; ok {
			delete(uniqueMap, n)
		} else {
			uniqueMap[n] = struct{}{}
		}
	}
	for k, _ := range uniqueMap {
		return k
	}
	return
}

func getUniqueNumberVersion02(numbers []int) int {
	k := 0
	for _, n := range numbers {
		k = k ^ n
	}
	return k
}
