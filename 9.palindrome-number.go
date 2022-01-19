package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	len := len(str)
	for i := 0; i < len/2; i++ {
		if str[i] != str[len-1-i] {
			return false
		}
	}

	return true
}

// @lc code=end
