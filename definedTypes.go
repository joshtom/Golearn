package main

import (
	"errors"
	"fmt"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

// Money is a defined type for monetary values in minor unites

type Money int64

type Price struct {
	Amount   Money
	Currency Currency
}

func (p Price) Add(other Price) (Price, error) {
	if p.Currency != other.Currency {
		return Price{}, errors.New("currency mismatch")
	}
	return Price{Amount: p.Amount + other.Amount, Currency: p.Currency}, nil
}

// String formats the price as human-readable.
func (p Price) String() string {
	return fmt.Sprintf("%s %.2f", p.Currency, float64(p.Amount)/100)
}

func definedTypes() {
	coffee := Price{Amount: 450, Currency: USD} // $4.50
	muffin := Price{Amount: 325, Currency: USD} // $3.25
	total, err := coffee.Add(muffin)

	if err != nil {
		panic(err)
	}
	fmt.Println("Total: ", total)

	// Compile-time safety: won’t let you add USD and EUR without handling it.
	lunchEU := Price{Amount: 1200, Currency: EUR}
	_, err = total.Add(lunchEU)
	fmt.Println("Cross-currency add: ", err)
}
