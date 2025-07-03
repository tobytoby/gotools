package main

import "fmt"

/*
 * 轮转数组
 * https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
 */

// a6B189TurnByCircle 环状替换
func a6B189TurnByCircle(nums []int, k int) {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	l := len(nums)
	k %= l           //防止k大于l
	cnt := gcd(k, l) //要跑的圈数
	for start := 0; start < cnt; start++ {
		//pre是当前要写入目标位置的值:初始为nums[start]
		//cur是当前指针位置
		pre, cur := nums[start], start

		//cur == start,则表示循环回到了初始位置
		for ok := true; ok; ok = cur != start {
			//计算目标位置
			next := (cur + k) % l
			//当前位置和目标位置的元素替换
			nums[next], pre = pre, nums[next]
			//把要循环的位置后一至下一个位置
			cur = next
		}
	}
}

// a6B189TurnReverse 数组反转
func a6B189TurnReverse(nums []int, k int) {
	var reverse func(nums []int)
	reverse = func(nums []int) {
		for i, n := 0, len(nums); i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	}

	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	a6B189TurnByCircle(nums, 8)
	fmt.Printf("num:%+v\n", nums)

	nums2 := []int{1, 2, 3, 4, 5, 6}
	a6B189TurnReverse(nums2, 8)
	fmt.Printf("num:%+v\n", nums2)
}
