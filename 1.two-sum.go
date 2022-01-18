package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var resX, resY int
Loop:
	for ix, x := range nums {
		for iy, y := range nums {
			if ix == iy {
				continue
			}

			if x+y == target {
				resX = ix
				resY = iy
				break Loop
			}
		}
	}

	return []int{resX, resY}
}

// @lc code=end
