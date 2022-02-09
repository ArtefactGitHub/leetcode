package leetcode

import "math"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	var max int = math.MinInt
	for i, v := range nums {
		for j := i; j >= 0; j-- {
			v = CalcSum(nums[j : i+1])
			if max < v {
				max = v
			}
		}
	}

	// max = 11081, time limit
	return max
}

func CalcSum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

// @lc code=end
