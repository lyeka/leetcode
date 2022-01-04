/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var sum int
	max_left := make([]int, len(height))
	max_right := make([]int, len(height))

	for i:=1; i<len(height); i++ {
		max_left[i] = max(max_left[i-1], height[i-1])
	}
	for i:=len(height)-2;i>0; i-- {
		max_right[i] = max(max_right[i+1], height[i+1])
	}
	for i:=1; i<len(height); i++ {
		m := min(max_left[i], max_right[i])
		if m > height[i] {
			sum += (m - height[i])
		}
		
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

