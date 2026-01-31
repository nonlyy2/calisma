package leetcode

// Задача №121: Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxDelta := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		delta := prices[i] - minPrice
		if delta > maxDelta {
			maxDelta = delta
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxDelta
}
