package main

import (
	"fmt"
)

/*
 * 买卖股票的最佳时机
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150
 */

type Profit struct {
	DayStart      int `json:"day_start"`
	DayStartPrice int `json:"day_start_price"`
	DayEnd        int `json:"day_end"`
	DayEndPrice   int `json:"day_end_price"`
	Profit        int `json:"profit"`
}

type BuySell struct {
	Day   int `json:"day"`
	Price int `json:"price"`
}

// a7b121BaoLi  暴力轮询法
func a7b121BaoLi(nums []int) *Profit {
	bs := new(Profit)
	l := len(nums)
	for x := 0; x < l-1; x++ {
		for y := x + 1; y < l; y++ {
			if nums[y] <= nums[x] {
				continue
			}
			if t := nums[y] - nums[x]; t > bs.Profit {
				bs = &Profit{
					DayStart:      x,
					DayStartPrice: nums[x],
					DayEnd:        y,
					DayEndPrice:   nums[y],
					Profit:        t,
				}
			}
		}
	}
	return bs
}

// a7b121V2 一次遍历, 用变量记录历史最低价和最高利润
func a7b121V2(nums []int) *Profit {
	//假设第一天买的
	buy := BuySell{
		Day:   0,
		Price: nums[0],
	}
	sell := BuySell{}
	profit := 0

	for i, v := range nums {
		if i == 0 {
			continue
		}
		if tmpProfit := v - buy.Price; tmpProfit > profit {
			sell = BuySell{
				Day:   i,
				Price: v,
			}
			profit = v - buy.Price
		}

		if v < buy.Price {
			buy = BuySell{
				Day:   i,
				Price: v,
			}
		}
	}
	if sell.Day == 0 {
		return nil
	}
	return &Profit{
		DayStart:      buy.Day,
		DayStartPrice: buy.Price,
		DayEnd:        sell.Day,
		DayEndPrice:   sell.Price,
		Profit:        profit,
	}
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4, 10, 8}
	fmt.Printf("%+v\n", a7b121BaoLi(nums))
	fmt.Printf("%+v\n", a7b121V2(nums))
}
