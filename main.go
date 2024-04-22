package main

import (
	"fmt"
	"wallet/pkg/bank/card"
	"wallet/pkg/bank/types"
)

func main() {
	visa := types.Card{
		ID:         1,
		PAN:        "5058333399994884",
		Balance:    9999,
		Currency:   types.TJS,
		Active:     true,
		Color:      "white",
		Name:       "Infinity",
		MinBalance: 3,
	}
	fmt.Println(visa.Balance)
	card.Withdraw(&visa, 1999)
	fmt.Println(visa.Balance)
}
