package main

import (
	"fmt"
)

/*
 * 接雨水
 * https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * a16b42Trap 动态规划
 * 对于下标i,下雨后雨水能达到的最大高度等于i两边高度的最小值,i处能接的雨水量等于i处水能达到的最大高度-height[i]
 * 扫描数组,记录左边和右边的最大高度,累加每个i能接的雨水量
 */
func a16b42Trap(height []int) (ans int) {
	n := len(height)
	if n <= 1 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}

	return
}

/*
 * a16b42Stack 单调栈，维护一个单调递减的栈
 * 从左到右遍历数组,遍历到下标i时,如果栈内至少有两个元素,记栈顶元素为top,top下面的一个元素是left,则一定有height[left] >= height[top].
 * 如果当前元素,height[i] > height[top]，则得到了一个可以蓄水的池子, 该区域的宽度是 w = i - left - 1, 高度是 h = min(height[i], height[left]) - height[top]
 *  w * h 即可以得到可蓄的雨水量。
 * 为了得到left,需要将top出栈。在对top计算能接的雨水量之后，left变成的新的top,如此重复,直到栈我空，或者栈顶对应的height >= height[i].
 * 在对下标i计算完能接的雨水量之后, 将i入栈。重复操作。
 */
func a16b42Stack(height []int) (ans int) {
	stack := make([]int, 0)
	for i, h := range height {
		//如果当前栈不为空,且当前位置高度大于栈顶元素高度:height[i] > height[top]
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			//将栈顶元素弹出
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curW := i - left - 1
			curH := min(h, height[left]) - height[top]
			ans += curW * curH
		}
		//将当前位置入栈
		stack = append(stack, i)
	}
	return
}

/*
 * a16b42Point 双指针
 */
func a16b42Point(height []int) (ans int) {
	n := len(height)
	left := 0
	right := n - 1
	leftMax := 0
	rightMax := 0
	//两个指针没有相遇
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

/*
 * a16b42FindHighest 找出最高的柱子
 * 柱子左侧水位一定是递增的，右侧水位一定是递减的，然后左侧蓄水和右侧蓄水累加
 */
func a16b42FindHighest(height []int) (ans int) {
	highest := 0
	highestPos := -1
	for i, h := range height {
		if h > highest {
			highest = h
			highestPos = i
		}
	}
	if highestPos == -1 {
		return 0
	}

	high := 0
	for i := 0; i < highestPos; i++ {
		//左侧水位
		if height[i] > high {
			high = height[i]
		}
		ans += high - height[i]
	}

	high = 0
	for i := len(height) - 1; i > highestPos; i-- {
		if height[i] > high {
			high = height[i]
		}
		ans += high - height[i]
	}
	return
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("%d\n", a16b42Stack(height))
	fmt.Printf("%d\n", a16b42Point(height))
	fmt.Printf("%d\n", a16b42FindHighest([]int{1, 3, 4, 5, 4, 4, 3, 1}))
}
