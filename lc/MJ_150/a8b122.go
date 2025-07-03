package main

import "fmt"

/*
 * 买卖股票的最佳时机:连日购买股票，当日可卖出,再买入
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * 动态规划
 *
 * dp[i][0]表示第i天交易完成后没有股票的收益
 *      本状态的转移状态为 前一天就没有持有股票:dp[i-1][0], 或者前一天持有股票dp[i-1][1],今天卖出
 *      那么今天的利润既为求两者的最大值: max(dp[i-1][0], dp[i-1][1]+price[i])
 * dp[i][1]表示第i天交易完成后持有股票的收益
 *      本状态的转移状态为  前一天就持有股票: dp[i-1][1], 或者前一天没有持有股票，今天买入
 *      那么今天的利润既为求两者的最大值: max(dp[i-1][1], dp[i-1][0]-price[i])
 * 基本常识: 最后一天交易结束之后，未持有股票的收益一定大于持有股票的收益
 */
func a8b122DP(prices []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	dp := make([][2]int, n)
	//第0天的收益很明确，未持有股票的收益为0，持有股票的收益为-price[0]
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	//最后一天交易结束之后，未持有股票的收益一定大于持有股票的收益
	return dp[n-1][0]
}

/*
 * a8b122DP的优化版本,因为每一天的状态只跟前一天的状态有关,跟更早的状态无关，所以我们只需要用两个变量存储前一天的状态即可
 */
func a8b122DPOpt(prices []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	d0 := 0
	d1 := -prices[0]
	for i, v := range prices {
		if i == 0 {
			continue
		}
		/*
		 * 注意这里不能是如下的代码，想想为什么
		 * d0 = max(d0, d1+v)
		 * d1 = max(d1, d0-v)
		 */
		d0, d1 = max(d0, d1+v), max(d1, d0-v)
	}

	return d0
}

/*
 * 贪心算法,只要有收益就算
 */
func a8b122Greedy(prices []int) int {
	profit := 0
	for i, v := range prices {
		if i == 0 {
			continue
		}
		if v > prices[i-1] {
			profit += v - prices[i-1]
		}
	}
	return profit
}

func main() {
	prices := []int{10, 6, 5, 20, 18, 5, 3, 17}
	fmt.Printf("profit:%d\n", a8b122DP(prices))
	fmt.Printf("profit:%d\n", a8b122DPOpt(prices))
	fmt.Printf("profit:%d\n", a8b122Greedy(prices))
}
