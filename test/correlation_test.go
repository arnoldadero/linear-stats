package test

import (
	"github.com/arnoldadero/linear-stats/regression"
	"math"
	"testing"
)

func TestCalculatePearsonCorrelation(t *testing.T) {
	testCases := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Perfect positive correlation",
			data:     []float64{1, 2, 3, 4, 5},
			expected: 1.0,
		},
		{
			name:     "Perfect negative correlation",
			data:     []float64{5, 4, 3, 2, 1},
			expected: -1.0,
		},
		{
			name:     "No correlation",
			data:     []float64{1, 1, 1, 1, 1},
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := regression.CalculatePearsonCorrelation(tc.data)
			if math.Abs(result-tc.expected) > 1e-6 {
				t.Errorf("Expected correlation: %f, got: %f", tc.expected, result)
			}
		})
	}
}
