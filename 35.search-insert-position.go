package leetcode

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	var f func(nums []int, head, tail, target int) int
	f = func(nums []int, head, tail, target int) int {
		cen := head + (tail-head)/2

		if tail < head {
			return cen
		}

		if nums[cen] == target {
			return cen
		} else if nums[cen] < target {
			head = cen + 1
		} else {
			tail = cen - 1
		}
		return f(nums, head, tail, target)
	}
	return f(nums, 0, len(nums)-1, target)
}

// @lc code=end
