package main

import "fmt"

/*
 * 跳跃游戏:判断是否可以调到最后一个位置
 * https://leetcode.cn/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * 贪心算法实现
 * 依次遍历数组中的每个元素，并且实时维护最远可以到达的位置
 * 对于当前遍历到的位置x，如果它在最远可以到达的范围内，那么我们可以从起点通过若干次跳跃到达该位置,因此我们可以用x+nums[x]更新最远可以到达的位置
 * 遍历过程中，如果最远可以到达的位置大于等于数组最后一个元素的位置，那说明可达
 */
func a9b55Greedy(locations []int) bool {
	if len(locations) <= 1 {
		return true
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	endLoc := len(locations) - 1
	rightMost := locations[0]
	for i, v := range locations {
		if i == 0 {
			continue
		}

		//当前位置是否可达
		if i > rightMost {
			return false
		}
		rightMost = max(rightMost, i+v)

		//最后一个位置是否可达
		if endLoc <= rightMost {
			return true
		}
	}
	return false
}

func main() {
	//可以到到达的位置
	a1 := []int{2, 3, 1, 1, 4}
	fmt.Printf("是否可以到达:%v\n", a9b55Greedy(a1))
	//不可以到到达的位置
	a2 := []int{3, 2, 1, 0, 4}
	fmt.Printf("是否可以到达:%v\n", a9b55Greedy(a2))
}
