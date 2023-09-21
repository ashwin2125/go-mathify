package arithmetic

import (
	"math"
	"testing"
)

func TestMultiplyInts(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 15},
		{3, 5, 15},
		{0, 0, 0},
		{-5, -3, 15},
		{math.MaxInt64, 1, math.MaxInt64},
		{math.MinInt64, 1, math.MinInt64},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiplyFloats(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{1.5, 1.0, 1.5},
		{1.0, 1.5, 1.5},
		{0.0, 0.0, 0.0},
		{-1.5, -1.0, 1.5},
		{math.Inf(1), 1, math.Inf(1)},
		{math.Inf(-1), 1, math.Inf(-1)},
		{math.NaN(), 1, math.NaN()},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}
			if got != tt.want {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
