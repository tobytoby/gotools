package string

/*
 * 字符串匹配KMP算法的实现,可以在不使用循环嵌套的情况下,在一个字符串中找到另外一个字符串的出现位置。
 * 常见场景有求子串首次出现的位置,以及子串出现的次数
 */

func KMPSearch(pattern string, text string) int {
	m := len(pattern)
	n := len(text)
	lps := make([]int, m)
	computeLPSArray(pattern, m, lps)

	i := 0
	j := 0
	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}
		if j == m {
			return i - j
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func computeLPSArray(pattern string, m int, lps []int) {
	len := 0
	lps[0] = 0
	i := 1
	for i < m {
		if pattern[i] == pattern[len] {
			len++
			lps[i] = len
			i++
		} else {
			if len != 0 {
				len = lps[len-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}
