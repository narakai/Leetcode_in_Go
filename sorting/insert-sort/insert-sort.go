package insert_sort

func insert(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := i
		temp := nums[i]
		for {
			if index <= 0 || nums[index-1] < temp {
				break
			}
			nums[index] = nums[index-1]
			index -= 1
		}
		//插入位置
		nums[index] = temp
	}
	return nums
}
