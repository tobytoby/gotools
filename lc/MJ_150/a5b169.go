package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
 * 多数元素
 * https://leetcode.cn/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150
 */

// a5b169FindModeV1 统计方法
func a5b169FindModeV1(nums []int) int {
	h := make(map[int]int, 0)
	cnt := 0
	mx := 0
	for _, v := range nums {
		h[v]++
		if h[v] >= cnt {
			cnt = h[v]
			mx = v
		}
	}
	return mx
}

// a5b169FindModeV2 排序法,因为众数的数量 >= n/2,所以下标n/2的数字一定是种属
func a5b169FindModeV2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// a5b169FindModeV3 随机法 因为大多数的数组下标都被众数占据了，所以我们随机挑选一个数字是否众数，大概率是能找到众数的
func a5b169FindModeV3(nums []int) int {
	l := len(nums)
	majorityCount := l / 2
	rand.NewSource(time.Now().Unix())
	for {
		candidate := nums[rand.Intn(l)]
		cnt := 0
		for _, v := range nums {
			if v == candidate {
				cnt++
			}
		}
		if cnt >= majorityCount {
			return candidate
		}
	}
}

// a5b169FindModeV4 使用分治思想,分成左右两部分,分别求出众数,再在两个众数中选出正确的众数
func a5b169FindModeV4(nums []int) int {
	var majorityElementRec func(lo, hi int) int
	majorityElementRec = func(lo, hi int) int {
		if lo == hi {
			return nums[lo]
		}

		mid := lo + (hi-lo)/2
		left := majorityElementRec(lo, mid)
		right := majorityElementRec(mid+1, hi)

		if left == right {
			return left
		}

		leftCnt, rightCnt := 0, 0
		for i := lo; i <= hi; i++ {
			if nums[i] == left {
				leftCnt++
			} else if nums[i] == right {
				rightCnt++
			}
		}
		if leftCnt > rightCnt {
			return left
		}
		return right
	}
	return majorityElementRec(0, len(nums)-1)
}

// a5b169FindModeV5 Boyer-Moore 投票算法,如果我们把众数记为 +1，把其他数记为 −1，将它们全部加起来，显然和大于 0，
// 从结果本身我们可以看出众数比其他数多。
// Boyer-Moore 算法的本质和方法四中的分治十分类似。 Boyer-Moore 算法的详细步骤：
// 我们维护一个候选众数 candidate 和它出现的次数 count。初始时 candidate 可以为任意值，count 为 0；
// 我们遍历数组 nums 中的所有元素，对于每个元素 x，在判断 x 之前，如果 count 的值为 0，我们先将 x 的值赋予 candidate，随后我们判断 x：
// 如果 x 与 candidate 相等，那么计数器 count 的值增加 1；
// 如果 x 与 candidate 不等，那么计数器 count 的值减少 1。
// 在遍历完成后，candidate 即为整个数组的众数。
func a5b169FindModeV5(nums []int) int {
	candidate, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 {
			candidate = v
		}
		if v == candidate {
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}

func main() {
	nums := []int{2, 5, 3, 2, 4, 5, 5, 6, 5, 10, 5, 5, 6, 5, 5, 10}
	fmt.Printf("num:%+v\n", a5b169FindModeV1(nums))
	fmt.Printf("num:%+v\n", a5b169FindModeV2(nums))
	fmt.Printf("num:%+v\n", a5b169FindModeV3(nums))
	fmt.Printf("num:%+v\n", a5b169FindModeV4(nums))
	fmt.Printf("num:%+v\n", a5b169FindModeV5(nums))
}
