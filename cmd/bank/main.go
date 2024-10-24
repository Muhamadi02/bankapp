package main

import (
	"bankapp/pkg/bank/card"
	"bankapp/pkg/bank/types"
	"fmt"
)

func main() {
	payCards := card.PaymentSources([]types.Card{
		{
			Name:    "card1",
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 999_99,
			Active:  true,
		},
		{
			Name:    "card2",
			PAN:     "5058 xxxx xxxx 7777",
			Balance: -500_00,
			Active:  true,
		},
	})

	fmt.Println(payCards)

	for _, tempCard := range payCards {
		fmt.Println(tempCard)
	}
}
