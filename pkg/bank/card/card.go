package card

import "bank/pkg/bank/types"

const withdrawalLimit = 200_000
const depositLimit = 50_000

func Issue(currency types.Currency, color, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "505800011110202",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

func Withdraw(card *types.Card, amount types.Money) {
	if amount <= 0 {
		return
	}
	if amount > withdrawalLimit {
		return
	}
	if !card.Active {
		return
	}
	if card.Balance < amount {
		return
	}
	card.Balance -= amount
}

func Deposit(card *types.Card, amount types.Money) {
	if amount <= 0 {
		return
	}
	if !card.Active {
		return
	}
	if amount > depositLimit {
		return
	}
	card.Balance += amount
}

func AddBonus(card *types.Card, percent, daysInMonth, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance < 1 {
		return
	}
	bonus := float64(card.MinBalance) * (float64(percent) / 100) * (float64(daysInMonth) / float64(daysInYear))
	card.Balance += types.Money(min(5000, bonus))
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var sources []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			sources = append(sources, types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			})
		}
	}
	return sources
}
