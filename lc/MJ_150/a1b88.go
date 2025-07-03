package main

import "fmt"

/*
 * 合并两个有序数组
 * https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
 */

// A1b88DoublePointerMerge 使用双指针,取两个队列头部最小的
func A1b88DoublePointerMerge(a1, a2 []int, m, n int) {
	mergeAry := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		//a1数组里的元素已经被拿完了
		if p1 == m {
			mergeAry = append(mergeAry, a2[p2:]...)
			break
		}

		if p2 == n {
			mergeAry = append(mergeAry, a1[p1:]...)
			break
		}

		if a1[p1] < a2[p2] {
			mergeAry = append(mergeAry, a1[p1])
			p1++
		} else {
			mergeAry = append(mergeAry, a2[p2])
			p2++
		}
	}
	copy(a1, mergeAry)
}

// A1b88ReverseDoublePointerMerge 还是使用双指针，不过是是使用逆向遍历的方式,取两个队列尾部最大的,放到a1数组的后半部分
func A1b88ReverseDoublePointerMerge(a1, a2 []int, m, n int) {
	p1, p2, tail := m-1, n-1, m+n-1
	for {
		if p1 == -1 {
			for p2 > -1 {
				a1[tail] = a2[p2]
				tail--
				p2--
			}
			break
		}

		if p2 == -1 {
			for p1 > -1 {
				a1[tail] = a1[p1]
				tail--
				p1--
			}
			break
		}

		if a1[p1] > a2[p2] {
			a1[tail] = a1[p1]
			p1--
			tail--
		} else {
			a1[tail] = a2[p2]
			p2--
			tail--
		}
	}
}

func main() {
	a1 := []int{1, 2, 3, 8, 0, 0, 0, 0}
	a2 := []int{2, 2, 4, 9}
	A1b88DoublePointerMerge(a1, a2, 4, 4)
	//A1b88ReverseDoublePointerMerge(a1, a2, 4, 4)
	fmt.Printf("a1:%+v\n", a1)
}
