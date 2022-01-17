/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package main
func isValid(s string) bool {
	d := map[string]string{
		"[":"]",
		"{":"}",
		"(":")",
	}
	t := []string{}
	for i:=0; i<len(s); i++ {
		k := string(s[i])
		v, ok := d[k]
		if ok {
			t = append(t, v)
		} else {
			if len(t) == 0 {
				return false
			}
			top := string(t[len(t)-1])
			if top != k {
				return false
			} else {
				t = t[:len(t)-1]
			}
		}
	}
	return len(t) == 0
}
// @lc code=end

