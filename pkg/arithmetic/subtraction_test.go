package arithmetic

import (
	"math"
	"testing"
)

func TestSubtractInts(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 2},
		{3, 5, -2},
		{0, 0, 0},
		{-5, -3, -2},
		{math.MaxInt64, 1, math.MaxInt64 - 1},
		{math.MinInt64, -1, math.MinInt64 + 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtractFloats(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{1.5, 1.0, 0.5},
		{1.0, 1.5, -0.5},
		{0.0, 0.0, 0.0},
		{-1.5, -1.0, -0.5},
		{math.Inf(1), 1, math.Inf(1)},
		{math.Inf(-1), -1, math.Inf(-1)},
		{math.NaN(), 1, math.NaN()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}
			if got != tt.want {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}