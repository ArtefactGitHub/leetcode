package leetcode

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	nlen := len(needle)
	if nlen == 0 {
		return 0
	}

	hlen := len(haystack)
	for i := 0; i < hlen; i++ {
		if haystack[i] == needle[0] {
			if hlen-i < nlen {
				return -1
			} else {
				if haystack[i:i+nlen] == needle {
					return i
				}
			}
		}
	}

	return -1
}

// @lc code=end
