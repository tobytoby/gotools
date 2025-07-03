package main

import (
	"sort"
)

/*
 * 求H指数，如果h有多个，取最大的那个
 * https://leetcode.cn/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150
 */

func a11b274Sort(nums []int) int {
	h := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0 && nums[i] > h; i-- {
		h++
	}
	return h
}

/*
 * 上个版本的时间复杂度，主要是跟排序有关,为了降低时间复杂度，我们可以使用计数排序算法，新建并维护一个counter数组用来记录当前
 * 引用次数的论文有几篇。
 * H指数不可能大于总发表的论文数量,所以对于引用次数超过论文发表数的情况,我们可以将其按照总的论文发表数记录。这样我们可以限制参与排序的
 * 数的大小在[0,n],使得计数排序的时间复杂度降低为O(n).
 * 从后向前遍历数组counter,对于每个0<=i<=n,在数组counter中得到大于或等于当前引用次数i的总论文数.当找到一个H指数时跳出循环.
 */
func a11b274SortOpt(nums []int) int {
	n := len(nums)
	//生成记录论文引用次数的计数器
	counter := make([]int, n+1)
	for _, v := range nums {
		if v >= n {
			counter[n]++
		} else {
			counter[v]++
		}
	}

	tot := 0
	for i := n; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}
	return 0
}

/*
 * 我们需要找到一个h值, 满足「有 h 篇论文的引用次数至少为 h」的最大值.小于等于h的所有值x都满足这个性质，而大于h的值都不满足这个性质.
 * 设查找范围的初始左边界left为0，初始右边界right为n.每次在查找范围内取中点mid,同时扫描整个数组,判断是否至少有mid个数大于mid.如果有，
 * 说明要找的h在搜索区间的右边，反之则在左边。
 */
func a11b274BinarySearch(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right + 1) >> 1
		cnt := 0
		for _, v := range nums {
			if v > mid {
				cnt++
			}
		}
		if cnt >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	//{0, 10, 30, 50, 60}
	nums := []int{30, 0, 60, 10, 50}
	print(a11b274Sort(nums))
	print(a11b274SortOpt(nums))
	print(a11b274BinarySearch(nums))
}
