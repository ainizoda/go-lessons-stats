package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	visa := card.Issue(types.USD, "black", "ginger")
	card.Deposit(&visa, 4000)
	fmt.Println(visa.Balance)
	card.Withdraw(&visa, 1999)
	fmt.Println(visa.Balance)
}
