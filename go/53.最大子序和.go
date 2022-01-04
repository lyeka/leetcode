/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res := nums[0]
	var sum int
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		res = max(res, sum)
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

