package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	var array []int
	for i, x := range nums {
		len := len(array)
		if i == 0 || array[len-1] != x {
			array = append(array, x)
		}
	}

	for i := 0; i < len(array); i++ {
		nums[i] = array[i]
	}

	return len(array)
}

// @lc code=end
