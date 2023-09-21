package arithmetic

import (
	"errors"
	"math"
	"testing"
)

func TestDivideInts(t *testing.T) {
	tests := []struct {
		a, b    int
		want    int
		wantErr error
	}{
		{10, 2, 5, nil},
		{-10, 2, -5, nil},
		{10, -2, -5, nil},
		{-10, -2, 5, nil},
		{0, 1, 0, nil},
		{0, -1, 0, nil},
		{1, 0, 0, errors.New("division by zero")},
		{math.MaxInt64, 1, math.MaxInt64, nil},
		{math.MinInt64, -1, math.MinInt64, nil},
		{math.MinInt64, 1, math.MinInt64, nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Divide() error = %v; wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivideFloats(t *testing.T) {
	tests := []struct {
		a, b, want float64
		wantErr    error
	}{
		{1.0, 2.0, 0.5, nil},
		{-1.0, 2.0, -0.5, nil},
		{1.0, -2.0, -0.5, nil},
		{-1.0, -2.0, 0.5, nil},
		{0.0, 1.0, 0.0, nil},
		{1.0, 0.0, math.Inf(1), nil},
		{-1.0, 0.0, math.Inf(-1), nil},
		{0.0, 0.0, math.NaN(), nil},
		{math.Inf(1), 1, math.Inf(1), nil},
		{math.Inf(-1), 1, math.Inf(-1), nil},
		{math.NaN(), 1, math.NaN(), nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Divide() error = %v; wantErr %v", err, tt.wantErr)
				return
			}
			if math.IsNaN(got) && math.IsNaN(tt.want) {
				return
			}
			if got != tt.want {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
