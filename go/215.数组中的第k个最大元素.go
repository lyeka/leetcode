/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	target := len(nums) - k 
	left, right := 0, len(nums)-1
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[target]
		}
		if index < target {
			left++
		}
		if index > target {
			right--
		}
	}
	return 0
}

func partition(nums []int, left, right int) int  {
	pivot := nums[left]
	j := left
	for i:=left; i<=right; i++ {
		if nums[i] < pivot {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}
// @lc code=end

//