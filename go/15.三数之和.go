/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
import "sort"
func threeSum(nums []int) [][]int {
	if len(nums) < 3  {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i:=0; i<len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r] 
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for (l<r && nums[l]==nums[l+1]) {
					l++
				}
				for (l<r && nums[r]==nums[r-1]) {
					r--
				}
				l++
				r--
			}
			if tmp > 0 {
				r--
			}
			if tmp < 0 {
				l++
			}
		}
		
	}
	return res
}
// @lc code=end

