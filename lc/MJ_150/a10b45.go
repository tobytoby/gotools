package main

import "fmt"

/*
 * 跳跃游戏2: 保证可以到达最后的位置，但是要使用最小的跳跃次数,且返回跳跃次数
 * https://leetcode.cn/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * 贪心算法: 反向查找出发位置
 * 考虑最后一步跳跃前所在的位置，该位置通过跳跃嫩到达最后一个位置。
 * 如果通过跳跃有多个位置可以到达最后一个位置，我们贪心的选择距离最后一个位置最远的那个位置,也就是对应下标最小的那个位置。
 * 因此我们可以从左到右遍历数组，选择第一个满足要求的位置，即就是距离最远的那个位置。
 * 找到最后异步跳跃之前的位置，再继续贪心的找倒数第二个位置，直到找到第0个位置
 * 疑问: 每一步都找最远的位置,由于上一步是最远，是否会导致下一步能找到的最远其实不是最远
 */
func a10b45Greedy(locations []int) int {
	//最后的位置
	endLoc := len(locations) - 1
	steps := 0
	for endLoc > 0 {
		curRoundFind := false
		for i := 0; i < endLoc; i++ {
			//如果最后一个位置能到
			if i+locations[i] >= endLoc {
				//继续下一个位置
				endLoc = i
				steps++
				curRoundFind = true
				break
			}
		}
		if !curRoundFind {
			return -1
		}
	}
	return steps
}

/*
 * 贪心算法2: 正向查找可到达的最远位置
 * 方法一时间复杂度高，如果我们正常查找每次可到达的最远位置,就可以在线性时间内得到最少的跳跃次数。
 * 所以只需要在当前可到大的位置找最远的即可
 * 具体实现: 维护当前能够到达的最大下标位置,记为边界.我们从左到右遍历数组,到达边界时,更新边界并将跳远次数+1
 *
 */
func a10b45Greedy2(locations []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(locations)
	end := 0
	maxLoc := 0
	steps := 0
	for i := 0; i < n-1; i++ {
		maxLoc = max(maxLoc, i+locations[i]) //反复更新 在跳跃范围内的元素能够提供的本次即将跳跃的最远距离
		if i == end {                        //遍历到上个起跳点能到的最远距离
			end = maxLoc //end更新为：右边界为i时候能提供的最远距离
			steps++      //先把步数+1，但实际还没跳（这也解释了循环终止条件）
		}
	}
	return steps
}

func main() {
	//可以到到达的位置
	a1 := []int{2, 3, 1, 1, 4}
	fmt.Printf("是否可以到达:%v\n", a10b45Greedy(a1))
	//不可以到到达的位置
	a2 := []int{3, 2, 1, 0, 4}
	fmt.Printf("是否可以到达:%v\n", a10b45Greedy(a2))
	//可以到到达的位置
	fmt.Printf("是否可以到达:%v\n", a10b45Greedy2(a1))
	//不可以到到达的位置
	fmt.Printf("是否可以到达:%v\n", a10b45Greedy2(a2))
}
