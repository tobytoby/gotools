package main

import "fmt"

/*
 * 分发糖果
 * https://leetcode.cn/problems/candy/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * a15b135V1 两次遍历
 * 我们可以将"相邻的孩子中，评分高的孩子必须获得更多的糖果"这句话拆分为两个规则：
 * 左规则: ratings[i-1] < ratings[i]时，i号学生的糖果数量将比i-1号孩子的糖果数量多
 * 右规则: ratings[i] > ratings[i+1]时, i号学生的糖果数量将比i+1号孩子的糖果数量多
 */
func a15b135V1(ratings []int) (ans int) {
	var max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	l := len(ratings)
	left := make([]int, l)
	//第一个孩子分1个
	left[0] = 1
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	/*
	 * 这里必须从最后一个孩子开始,不能用以下的代码
	 * right := 1
	 * for i := l - 2; i >= 0; i-- {
	 *		if ratings[i] > ratings[i+1] {
	 * 如果将最后一个孩子初始化为1,从倒数第二个开始遍历,则加糖果的时候，会漏掉最后一个孩子，如果是这样,则需要手动加上最后一个孩子的糖果:
	 * ans += max(left[l-1], 1)
	 */
	for i := l - 1; i >= 0; i-- {
		if i < l-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		//这里取max,不是min的原因是,当前学生的糖果数量必须要满足左右规则,那就只能取其中大的那个
		ans += max(left[i], right)
	}
	return
}

/*
 * a15b135V2 常数空间遍历，糖果数量总是尽量少给，且从1开始累计，要么比相邻的孩子多给1个，要么重置为1. 规律如下：
 * 从左到右枚举每一个同学，记前一个同学分得的糖果数量为 pre：
 *   如果当前同学比上一个同学评分高，说明我们就在最近的递增序列中，直接分配给该同学 pre+1 个糖果即可
 *   否则我们就在一个递减序列中，我们直接分配给当前同学一个糖果，并把该同学所在的递减序列中所有的同学都再多分配一个糖果，以保证糖果数量还是满足条件
 *       我们无需显式地额外分配糖果，只需要记录当前的递减序列长度，即可知道需要额外分配的糖果数量
 *       同时注意当当前的递减序列长度和上一个递增序列等长时，需要把最近的递增序列的最后一个同学也并进递减序列中
 *   这样，我们只要记录当前递减序列的长度 dec，最近的递增序列的长度 inc 和前一个同学分得的糖果数量 pre 即可
 */
func a15b135V2(ratings []int) (ans int) {
	n := len(ratings)
	ans = 1  //总分发糖果数量
	inc := 1 //最近的递增序列长度
	dec := 0 //当前递减序列长度
	pre := 1 //前一个孩子的糖果数量
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}

func main() {
	//ratings := []int{1, 0, 2} //5
	ratings := []int{1, 2, 2} //4
	fmt.Printf("v1:%d\n", a15b135V2(ratings))
}
