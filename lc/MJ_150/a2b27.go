package main

import (
	"fmt"
)

/*
 * 移除数组元素
 * https://leetcode.cn/problems/remove-element/solutions/730203/yi-chu-yuan-su-by-leetcode-solution-svxi/?envType=study-plan-v2&envId=top-interview-150
 */

// A2b27RemoveSpecialEle 实现逻辑:因为是要删除元素,所以输出数组的长度一定是小于等输入数组的,所以可以将输出数组保存在输入数组中。
// 可以使用双指针，右指针 right 指向下一个要处理的元素,左指针 left 指向下一个要赋值的位置.
// 如果右指针指向的元素不等于要删除的元素，它一定是输出元素,则right 和 left 都右移
// 如果右指针指向的元素等于要删除的元素，那么它不能在输出数组里，left不移动， right右移
func A2b27RemoveSpecialEle(nums []int, target int) int {
	//左指针
	left := 0
	//遍历相当于是直接的右指针移动
	for _, v := range nums {
		if v != target {
			nums[left] = v
			left++
		}
	}
	return left
}

// A2b27RemoveSpecialEleV2 优化项: 如果初始数组是[1,2,3,4,5],给定值为1,那么则需要将后面四个元素都左移动.实际上只需要将数组末尾的元素和当前值替换
func A2b27RemoveSpecialEleV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		if left == right {
			if nums[left] != target {
				left++
			}
			break
		}

		if nums[left] == target {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}

// A2b27RemoveSpecialEleV3 上面函数也可以简化为
func A2b27RemoveSpecialEleV3(nums []int, target int) int {
	//疑问这里右指针，为什么不是 len(nums) - 1
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == target {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	nums := []int{3, 2, 2, 3, 5}
	k := A2b27RemoveSpecialEleV3(nums, 2)
	fmt.Printf("num:%+v\n", nums)
	fmt.Printf("k:%d\n", k)
}
