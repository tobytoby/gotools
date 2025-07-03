package main

import "fmt"

/*
 * 删除有序数组中的重复项
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
 */

// A3B26DelRepeat 因为是有序数组,所以相同的元素一定是连续的,可以使用快 fast 慢 slow 指针
// 快指针表示遍历数组到达的下标位置,慢指针表示下一个不同元素要填入的下标位置,初始时的两个指针都指向下标1.
// 假设数组长度为n.将快指针依次从1到n-1的每个位置,对于每个位置，如果nums[fast] != nums[fast-1],说明nums[fast]和之前的元素不同
// 因此将nums[fast]的复制到nums[slow],然后将slow的值右移
// 遍历结束之后，从nums[0]到nums[slow-1]的每个元素都是唯一的
func A3B26DelRepeat(nums []int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	slow, fast := 1, 1
	for fast < l {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := A3B26DelRepeat(nums)
	fmt.Printf("num:%+v\n", nums)
	fmt.Printf("k:%d\n", k)
}
