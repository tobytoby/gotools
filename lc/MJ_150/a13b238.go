package main

import "fmt"

/*
 * 除自身外数组的乘积,必须在O(n)时间复杂度完成,且不能使用除法
 * https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * a13B238JiSumByDiv 先用除法试试
 * 但是如果数组中有0会失效,所以需要处理0的情况
 * 如果数组中有一个0,那么在0的所在的位置，有值,其他位置都是0，如果有超过两个以上的0,那么其余位置全部是0
 */
func a13B238JiSumByDiv(nums []int) []int {
	l := len(nums)
	sum := 1
	zeroNum := 0
	zeroIdx := -1
	for i, v := range nums {
		if v == 0 {
			zeroIdx = i
			zeroNum++
			if zeroNum >= 2 {
				break
			}
			continue
		}
		sum *= v
	}
	switch zeroNum {
	case 0:
		for i, v := range nums {
			nums[i] = sum / v
		}
		return nums
	case 1:
		nums = make([]int, l)
		nums[zeroIdx] = sum
		return nums
	default:
		return make([]int, l)
	}
}

/*
 * a13B238JiSumByFillLR 填充左右乘积列表, 当前元素只需要左右乘积列表相乘即可
 */
func a13B238JiSumByFillLR(nums []int) []int {
	length := len(nums)
	//l, r 分别表示左右两侧的乘积列表
	l, r, answer := make([]int, length), make([]int, length), make([]int, length)

	l[0] = 1
	for i := 1; i < length; i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	r[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	for i := 0; i < length; i++ {
		answer[i] = l[i] * r[i]
	}
	return answer
}

/*
 * a13B238JiSumByFillLROpt  上面的函数空间复杂度高,可以只填充左边
 */
func a13B238JiSumByFillLOpt(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	//将answer[i]先作为左侧所有元素的乘积
	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	r := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		answer[i] = answer[i] * r
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		r *= nums[i]
	}
	return answer
}

/*
 * a13B238JiSumByFillLROpt  上面的函数空间复杂度高,可以只填充右边
 */
func a13B238JiSumByFillROpt(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	answer[length-1] = 1
	//将answer[i]先作为右侧所有元素的乘积
	for i := length - 2; i >= 0; i-- {
		answer[i] = answer[i+1] * nums[i+1]
	}

	l := 1
	for i := 0; i < length; i++ {
		// 对于索引 i，右边的乘积为 answer[i]，左边的乘积为 l
		answer[i] = answer[i] * l
		// l 需要包含左边所有的乘积，所以计算下一个结果时需要将当前值乘到 l 上
		l *= nums[i]
	}
	return answer
}

func main() {
	nums1 := []int{2, 3, 4, 6}
	nums2 := []int{2, 3, 0, 6, 7}
	nums3 := []int{2, 3, 0, 6, 0, 6}

	fmt.Printf("v2:%+v\n", a13B238JiSumByFillLR(nums1))
	fmt.Printf("v2:%+v\n", a13B238JiSumByFillLR(nums2))
	fmt.Printf("v2:%+v\n", a13B238JiSumByFillLR(nums3))

	fmt.Printf("v2:%+v\n", a13B238JiSumByFillLOpt(nums1))
	fmt.Printf("v2:%+v\n", a13B238JiSumByFillLOpt(nums2))
	fmt.Printf("v2:%+v\n", a13B238JiSumByFillLOpt(nums3))

	fmt.Printf("v2:%+v\n", a13B238JiSumByFillROpt(nums1))
	fmt.Printf("v2:%+v\n", a13B238JiSumByFillROpt(nums2))
	fmt.Printf("v2:%+v\n", a13B238JiSumByFillROpt(nums3))
}
