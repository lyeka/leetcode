/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	set := make(map[byte]int)
	res := 1
	l, r := 0,1
	set[s[0]] = 1
	for r < len(s) {
		if set[s[r]] == 0 {
			set[s[r]]=1
			res = max(res, r-l+1)
			r++		
		} else {
			set[s[l]]=0
			l++
		}
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

