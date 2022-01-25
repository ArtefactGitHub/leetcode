package leetcode

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var root *ListNode
	if list1 != nil || list2 != nil {
		result = &ListNode{}
		root = result
	}
	// どちらかにまだ要素がある場合ループ処理を継続させる
	for list1 != nil || list2 != nil {
		// どちらもnilではない場合、Valの小さい方を追加する
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				result.Val = list1.Val
				list1 = list1.Next
			} else {
				result.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 == nil {
			result.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			result.Val = list1.Val
			list1 = list1.Next
		}

		if list1 != nil || list2 != nil {
			next := ListNode{}
			result.Next = &next
			result = &next
		}
	}

	return root
}

// @lc code=end
