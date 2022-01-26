package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	var counter int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[counter-1] {
			nums[counter] = nums[i]
			counter++
		}
	}

	return counter
}

// @lc code=end
