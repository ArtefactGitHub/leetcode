package leetcode

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	var counter int
	max := len(strs[0])
LOOP:
	for index := 0; index < max; index++ {
		var b byte
		for i, str := range strs {
			if i == 0 {
				b = str[index]
			} else if len(str) <= index || b != str[index] {
				break LOOP
			}
		}

		counter++
	}

	return strs[0][:counter]
}

// @lc code=end
