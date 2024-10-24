package card

import (
	"bankapp/pkg/bank/types"
	"fmt"
)

func ExampleCard() {
	card := types.Card{Balance: 20_000_00, Active: true}
	clone := card
	card.Balance -= 10_000_00
	fmt.Println(clone.Balance)
	// Output: 2000000
}

func ExampleWithdraw() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}
func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 15_000_00, Active: true}
	Withdraw(&card, 20_000_00)
	fmt.Println(card.Balance)
	// Output: 1500000
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Active)
	// Output: false
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Withdraw(&card, 21_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Active)
	// Output: false
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 30_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 20_000, MinBalance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 20024
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000, MinBalance: 10_000, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 20000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -1_000, MinBalance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -1000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 20_000, MinBalance: 3_000_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 20000

}

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))

	// Output:
	// 100000
	// 300000

}

func ExamplePaymentSources() {
	payCards := PaymentSources([]types.Card{
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
		{
			Name:    "card3",
			PAN:     "5058 xxxx xxxx 2222",
			Balance: -500_00,
			Active:  false,
		},
		{
			Name:    "card4",
			PAN:     "5058 xxxx xxxx 3333",
			Balance: 900_00,
			Active:  false,
		},
		{
			Name:    "card5",
			PAN:     "5058 xxxx xxxx 1111",
			Balance: 1_00,
			Active:  true,
		},
	})

	for _, card := range payCards {
		fmt.Println(card.Number)
	}

	// Output:
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 1111
}
