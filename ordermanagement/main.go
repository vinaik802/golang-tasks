package main

import (
	"fmt"
	"strings"
)

type IMarketLimit interface {
	Execute(marketprice float64) error
}
type MarketOrder struct {
	qty   int
	scrip string
}
type LimitOrder struct {
	qty        int
	limitPrice float64
	scrip      string
}

func main() {
	marketprice := 51.23
	var scrip string
	var price float64
	var qty int
	var Otype string
	println("Enter the order type : m for market order and l for limit order")
	fmt.Scanln(&Otype)
	if strings.ToLower(Otype) == "l" {
		fmt.Println("Enter the scrip")

		fmt.Scanln(&scrip)

		fmt.Println("Enter the price")
		fmt.Scanln(&price)

		fmt.Println("Enter the quantity")
		fmt.Scanln(&qty)
	} else if strings.ToLower(Otype) == "m" {
		fmt.Println("Enter the scrip")

		fmt.Scanln(&scrip)

		fmt.Println("Enter the quantity")
		fmt.Scanln(&qty)
	} else {
		panic("entered wrong input")
	}

	marketO := newMarket(qty, scrip)
	limitO := newLimit(qty, price, scrip)

	if strings.ToLower(Otype) == "market" {
		if err := marketO.Execute(marketprice); err != nil {

			println(err.Error())
		}
	} else {
		if err := limitO.Execute(marketprice); err != nil {
			println(err.Error())
		}
	}

}

func newMarket(qty int, scrip string) IMarketLimit {
	return &MarketOrder{qty, scrip}
}

func newLimit(qty int, limitprice float64, scrip string) IMarketLimit {
	return &LimitOrder{qty, limitprice, scrip}
}

func (m *MarketOrder) Execute(marketprice float64) error {
	fmt.Println("Processing Market Order: Buying ", m.scrip, " at  Price(", marketprice, ")")
	return nil

}

func (l *LimitOrder) Execute(marketprice float64) error {
	if l.limitPrice < float64(marketprice) {
		return fmt.Errorf("limit price is below than market price")
	}
	fmt.Println("processing Limit Order: Buying ", l.qty, " ", l.scrip, " at Market Price(", l.limitPrice, ")")
	return nil
}
