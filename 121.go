package main

func maxProfit(prices []int) int {
	var minimumPrice, maxProfit = prices[0], 0

	for i := 1; i < len(prices); i++ {
		minimumPrice = min(minimumPrice, prices[i])
		maxProfit = max(maxProfit, prices[i] - minimumPrice)
	}

	return maxProfit
}
