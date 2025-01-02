package jct

import (
	"time"

	tz "github.com/mi-wada/jct-go/internal"
)

// Rate returns the tax rate applicable at the specified time.
// If the time is before the earliest known rate, it returns 0.0.
func Rate(at time.Time) float64 {
	for i := len(rates) - 1; i >= 0; i-- {
		if at.Equal(rates[i].startDate) || at.After(rates[i].startDate) {
			return rates[i].rate
		}
	}

	return 0.0
}

// Tax returns the tax amount for the given amount at the specified time.
func Tax(amount float64, at time.Time) float64 {
	return amount * Rate(at)
}

// Total returns the total amount (including tax) for the given amount at the specified time.
func Total(amount float64, at time.Time) float64 {
	return amount + Tax(amount, at)
}

type rate struct {
	startDate time.Time
	rate      float64
}

var rates = []rate{
	{
		startDate: time.Date(1989, 4, 1, 0, 0, 0, 0, tz.JST),
		rate:      0.03,
	},
	{
		startDate: time.Date(1997, 4, 1, 0, 0, 0, 0, tz.JST),
		rate:      0.05,
	},
	{
		startDate: time.Date(2014, 4, 1, 0, 0, 0, 0, tz.JST),
		rate:      0.08,
	},
	{
		startDate: time.Date(2019, 10, 1, 0, 0, 0, 0, tz.JST),
		rate:      0.10,
	},
}
