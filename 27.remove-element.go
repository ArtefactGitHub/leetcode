package leetcode

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	var counter int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}
	}

	return counter
}

// @lc code=end
