package jct_test

import (
	"testing"
	"time"

	"github.com/mi-wada/jct-go"
	tz "github.com/mi-wada/jct-go/internal"
)

func TestRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		at   time.Time
		want float64
	}{
		{
			name: "No consumption tax era",
			at:   time.Date(1989, 4, 1, 0, 0, 0, 0, tz.JST).Add(-1 * time.Nanosecond),
			want: 0.00,
		},
		{
			name: "Start of 0.03",
			at:   time.Date(1989, 4, 1, 0, 0, 0, 0, tz.JST),
			want: 0.03,
		},
		{
			name: "End of 0.03",
			at:   time.Date(1997, 4, 1, 0, 0, 0, 0, tz.JST).Add(-1 * time.Nanosecond),
			want: 0.03,
		},
		{
			name: "Start of 0.05",
			at:   time.Date(1997, 4, 1, 0, 0, 0, 0, tz.JST),
			want: 0.05,
		},
		{
			name: "End of 0.05",
			at:   time.Date(2014, 4, 1, 0, 0, 0, 0, tz.JST).Add(-1 * time.Nanosecond),
			want: 0.05,
		},
		{
			name: "Start of 0.08",
			at:   time.Date(2014, 4, 1, 0, 0, 0, 0, tz.JST),
			want: 0.08,
		},
		{
			name: "End of 0.08",
			at:   time.Date(2019, 10, 1, 0, 0, 0, 0, tz.JST).Add(-1 * time.Nanosecond),
			want: 0.08,
		},
		{
			name: "Start of 0.10",
			at:   time.Date(2019, 10, 1, 0, 0, 0, 0, tz.JST),
			want: 0.10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := jct.Rate(tt.at)
			if got != tt.want {
				t.Errorf("Rate(%v) = %v, want %v", tt.at, got, tt.want)
			}
		})
	}
}