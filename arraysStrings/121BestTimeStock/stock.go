package besttimestock

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func MaxProfit(prices []int) int {
    // Handle edge case for empty input
    if len(prices) == 0 {
        return 0
    }

    // Initialize tracking variables
    minPrice := prices[0]
    maxProfit := 0

    for _, price := range prices {
        // Update max profit if selling at the current price is better
        if price-minPrice > maxProfit {
            maxProfit = price - minPrice
        }

        // Update the minimum price if the current price is lower
        if price < minPrice {
            minPrice = price
        }
    }

    return maxProfit
}