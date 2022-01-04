/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return 
	}

	index := partition(nums, left, right)
	quickSort(nums, left, index-1)
	quickSort(nums, index+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	j := left
	for i:=left+1; i<=right; i++ {
		if nums[i] < pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}
// @lc code=end

