package main

import "fmt"

/*
 * 删除有序数组中的重复项
 * 重复元素最多出现两次
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
 */

// A4B80DelRepeatN 我们需要检查的是上上个应该被保留的元素nums[slow-2] != nums[fast]
func A4B80DelRepeatN(nums []int, n int) int {
	l := len(nums)
	if l <= 2 {
		return n
	}
	slow, fast := n, n
	for fast < l {
		if nums[slow-n] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := A4B80DelRepeatN(nums, 2)
	fmt.Printf("num:%+v\n", nums)
	fmt.Printf("k:%d\n", k)
}
