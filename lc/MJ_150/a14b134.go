package main

import (
	"fmt"
	"log"
)

/*
 * 加油站: 环形加油站， 第i个加油站有汽油gas(i)升
 */

// 暴力遍历a14b134FindStart
func a14b134FindStart(gas, cost []int) int {
	step := 0
	l := len(gas)
	leftGas := 0
	start := 0
	startCp := 0
	hasTryStart := 1
LABEL:
	leftGas = 0
	for {
		leftGas += gas[start]
		//退出规则是判断当前邮箱的油是否能到达下一个地方
		if leftGas < cost[start] {
			//如果不能走下去，尝试从下一个出发点出发
			//首先要判断是不是所有出发点都尝试过了
			if hasTryStart >= l {
				return -1
			}

			startCp = (startCp + 1) % l
			start = startCp
			hasTryStart++
			goto LABEL
		}
		log.Printf("%+v\n", map[string]any{
			"开始位置":      startCp,
			"现在走到":      start,
			"剩余汽油":      leftGas,
			"走至下一步需要消耗": cost[start],
			"已经尝试次数":    hasTryStart,
		})

		//判断是否即将走完一圈
		if step >= l-1 {
			return startCp
		}

		leftGas -= cost[start]
		start = (start + 1) % l
		step++
	}
}

/*
 * a14b134FindStartGreedy 贪心算法实现
 * 如果 gas 总和 小于 cost 总和，无论从哪里出发都不可能绕一圈 → 直接返回 -1
 * 否则一定有解，从头开始遍历，记录每一段的剩余油量 tank
 * 如果tank < 0, 表示当前起点不行，换下一个
 */
func a14b134FindStartGreedy(gas, cost []int) int {
	totalTank := 0 // 总油量（用来判断是否有解）
	currTank := 0  // 当前从 start 到 i 的油量
	start := 0     // 起始加油站索引

	for i := 0; i < len(gas); i++ {
		gain := gas[i] - cost[i]
		totalTank += gain
		currTank += gain

		// 如果当前总油量小于0，说明无法从 start 到达 i+1，重新选择起点
		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}

	// 如果总油量大于等于0，说明肯定有解，返回起点
	if totalTank >= 0 {
		return start
	} else {
		return -1
	}
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	//gas := []int{2, 3, 4}
	//cost := []int{3, 4, 3}
	fmt.Printf("v1:%d\n", a14b134FindStart(gas, cost))
	fmt.Printf("v1:%d\n", a14b134FindStartGreedy(gas, cost))
}
