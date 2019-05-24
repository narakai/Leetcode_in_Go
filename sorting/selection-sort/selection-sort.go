package selection_sort

import "fmt"

//不断地选择 {剩余} 元素中的最小者
func selection(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		//每次循环重置minIndex
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		//在内层循环结束，也就是找到本轮循环的最小的数以后，再进行交换
		print(minIndex)
		print(i)
		fmt.Printf("before————%v", nums)
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
		fmt.Printf("after————%v", nums)
		print("\n")
	}
	return nums
}
