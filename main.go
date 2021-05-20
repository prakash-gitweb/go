package main

import (
	"fmt"
	"github.com/prakash-gitweb/go/lib"
)

type order struct {
	ticket  uint
	volume  float32
	symbol  string
	account uint
	price   float32
}

func main() {
	closed := order{
		ticket:  100000,
		volume:  100,
		symbol:  "EURUSD.ecn",
		account: 444499,
		price:   1.120098,
	}
	execute := lib.IbDeposit(&closed.ticket, &closed.volume, &closed.symbol, &closed.account)
	fmt.Println(execute)
}
