package card

import (
	"bankapp/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {

	card := types.Card{
		ID:      1000,
		PAN:     "5058 xxxx xxxx 0001",
		Balance: 0,
		Active:  true,
	}

	card.Currency = currency
	card.Color = color
	card.Name = name
	return card
}

const withdrawLimit = 20_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if card.Balance < amount {
		return
	}

	if amount < 0 {
		return
	}

	if amount > withdrawLimit {
		return
	}

	card.Balance -= amount

}

const depositLimit = 50_000_00

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > depositLimit {
		return
	}

	card.Balance += amount
}

const bonusLimit = 5_000

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance < 0 {
		return
	}

	bonus := types.Money(int(card.MinBalance)*percent*daysInMonth/daysInYear) / 100

	if bonus > bonusLimit {
		return
	}

	card.Balance += bonus
}

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		sum += card.Balance
	}
	return sum
}
