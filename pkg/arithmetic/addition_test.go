package arithmetic

import (
	"math"
	"testing"
)

func TestAddInts(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 1, 2},
		{-1, 1, 0},
		{100, 200, 300},
		{math.MaxInt64, 1, math.MinInt64},  // overflow
		{math.MinInt64, -1, math.MaxInt64}, // underflow
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestAddFloats(t *testing.T) {
	tests := []struct {
		a, b float64
		want float64
	}{
		{1.0, 1.0, 2.0},
		{-1.5, 1.5, 0.0},
		{100.25, 200.75, 301.0},
		{math.Inf(1), 1, math.Inf(1)},    // positive infinity
		{math.Inf(-1), -1, math.Inf(-1)}, // negative infinity
		{math.NaN(), 1, math.NaN()},      // not a number
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
