package main

import "fmt"

```
Maximum profit with K transactions

You're given an array of integers representing the prices of a single stock on various days. Each index in the array represents a different day. You're also given an integer 'k' which represents the number of transactions you're allowed to make. One transaction consists of buying the stock on a given day and selling it on another, later day.

Write a function that returns the maximum profit that you can make by executing 'k' transactions.

Note: You can't buy more than one share of the stock on the same day. You can't buy a share if 
you're already holding one. You may not use all the 'k' transactions you're allowed.

Example: prices = {5, 11, 3, 50, 60, 90}
K = 2

Output: 93
Method: 
Transaction 1: Buy 5, Sell 11
Transaction 2: Buy 3, Sell 90
```

type transactions struct {
	day       int
	buyPrice  int
	sellPrice int
}

func main() {
	stockPrices := []int{5, 11, 3, 50, 60, 90}
	maxTransactions := 2

	// TODO: validate data points

	// 1. store transaction history
	// day, buy price & sell price
	history := make([]*transactions, 0)

	// 2. logic to decide when to sell??
	// -> 5, 3, 30, 50, 60, 90
	// if i reach till final day its last??? we can go ahead and calculate
	// the profit

	// 3. calculate the profit

	lastTransDay := 0
	// i can make only upto 2 transactions (buys)
	for i := 0; i < maxTransactions; i++ {
		// buy the first value and decide sell and signal to continue
		buyDay := lastTransDay
		if lastTransDay == len(stockPrices) {
			buyDay = lastTransDay - 1
		}
		trans := &transactions{}
		trans.day = buyDay

		trans.buyPrice = stockPrices[buyDay] //index
		// sell price??
		sellDay, sellPrice := sellPrice(lastTransDay, trans.buyPrice, stockPrices)
		fmt.Println("sellDay, sellPrice", sellDay, sellPrice)
		trans.sellPrice = sellPrice
		history = append(history, trans)
		lastTransDay = sellDay
		if sellDay == len(stockPrices) {
			break
		}
	}

	// cal profit
	profit := 0
	for _, h := range history {
		fmt.Println(h)
		profit += (h.sellPrice - h.buyPrice)
	}

	fmt.Println(profit)
}

// what if buy & sell last day
// sell logic was not clear!!
// signal for last
func sellPrice(buyDay, buyPrice int, stockPrices []int) (int, int) {
	stockPricesSl := stockPrices[buyDay:]

	lastSellDay := 0
	lastSellPrice := buyPrice
	for sellDay, sellPrice := range stockPricesSl {
		// find max sell price possible
		if lastSellPrice < sellPrice {
			lastSellPrice = sellPrice
			lastSellDay = sellDay
		}
		// if price is less than previos day price and less than buy price
		if sellPrice < lastSellPrice && sellPrice < buyPrice {
			// fmt.Println("sellDay, sellPrice", sellDay, sellPrice)
			// fmt.Println("lastsellDay, lastsellPrice", lastSellDay, lastSellPrice)
			// use last best price
			break
		}
	}

	return lastSellDay + (buyDay + 1), lastSellPrice
}
