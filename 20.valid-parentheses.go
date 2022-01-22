package leetcode

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	var stack []string
	for _, b := range s {
		c := string(b)

		// 開き括弧の場合は閉じ括弧をスタックに積む
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack, pairOfParentheses(c))
		} else {
			// 閉じの場合はスタックと同じか検証
			len := len(stack)

			// 最初に閉じ括弧が与えられた場合、またはスタックと異なる場合は不正な入力
			if len == 0 || stack[len-1] != c {
				return false
			} else {
				// 同じ場合はスタックを減らす
				stack = stack[:len-1]
			}
		}
	}

	// スタックが解消されていない場合は不正な入力
	if len(stack) == 0 {
		return true
	}
	return false
}

func pairOfParentheses(c string) string {
	switch c {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	default:
		return ""
	}
}

// @lc code=end
