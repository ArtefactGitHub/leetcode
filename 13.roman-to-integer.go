package leetcode

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	var result int
	len := len(s)
	for i := 0; i < len; i++ {
		if i == len-1 {
			result += toInt(s[i])
			break
		}

		now := toInt(s[i])
		next := toInt(s[i+1])
		// s[i] < s[i+1]の場合
		if now < next {
			// s[i+1] - s[i]を総和に加算する
			result += next - now
			// i をインクリメントする
			i++
		} else {
			// s[i] >= s[i+1]の場合
			// s[i]を総和に加算する
			result += now
		}
	}

	return result
}

func toInt(c byte) int {
	s := string(c)
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

// @lc code=end
