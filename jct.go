// Package jct provides functions to calculate the Japanese Consumption Tax.
package jct

import (
	"math/big"
	"time"

	"github.com/mi-wada/jct-go/internal/tz"
)

// Rate returns the tax rate applicable at the specified time.
// If the time is before the earliest known rate, it returns big.NewRat(0, 1).
func Rate(at time.Time) *big.Rat {
	for i := len(rates) - 1; i >= 0; i-- {
		if at.Equal(rates[i].from) || at.After(rates[i].from) {
			return rates[i].rate
		}
	}

	return big.NewRat(0, 1)
}

// Tax returns the tax amount for the given amount at the specified time.
// The tax amount is rounded down to the nearest integer.
func Tax(amount int64, at time.Time) int64 {
	amountRat := new(big.Rat).SetInt64(amount)
	r := Rate(at)

	taxRat := new(big.Rat).Mul(amountRat, r)
	taxBigint := new(big.Int).Div(taxRat.Num(), taxRat.Denom())

	return taxBigint.Int64()
}

// Total returns the total amount (including tax) for the given amount at the specified time.
// The total amount is rounded down to the nearest integer.
func Total(amount int64, at time.Time) int64 {
	return amount + Tax(amount, at)
}

type rate struct {
	from time.Time
	rate *big.Rat
}

var rates = []rate{
	{
		from: time.Date(1989, 4, 1, 0, 0, 0, 0, tz.JST),
		rate: big.NewRat(3, 100),
	},
	{
		from: time.Date(1997, 4, 1, 0, 0, 0, 0, tz.JST),
		rate: big.NewRat(5, 100),
	},
	{
		from: time.Date(2014, 4, 1, 0, 0, 0, 0, tz.JST),
		rate: big.NewRat(8, 100),
	},
	{
		from: time.Date(2019, 10, 1, 0, 0, 0, 0, tz.JST),
		rate: big.NewRat(10, 100),
	},
}
