package string

/*
 * 字符串匹配,朴素匹配算法,就是直接遍历循环,查找是否匹配
 */

func SimpleSearch(pattern string, text string) int {
	ptL := len(pattern)
	ttL := len(text)

	y := 0
	start := -1
	for i := 0; i < ptL; i++ {
		if i >= ptL || y >= ttL {
			break
		}
		for ; y < ttL; y++ {
			if pattern[i] == text[y] {
				if i == 0 {
					start = y
				}
				y += 1
				break
			} else {
				start = -1
				i = 0
			}
		}
	}
	return start
}

/*
 * 通过快慢指针进行实现,只能判断是否包含,不能判断连续性
 */

func DoublePointerSearch(pattern, text string) int {
	i, j := 0, 0
	for i < len(pattern) && j < len(text) {
		if pattern[i] == text[j] {
			i++
		}
		j++
	}
	if i == len(pattern) {
		return j - len(pattern)
	}
	return -1
}
