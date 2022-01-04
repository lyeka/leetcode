/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i-1] + nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		res = max(res, nums[i])
	}

	return res

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

