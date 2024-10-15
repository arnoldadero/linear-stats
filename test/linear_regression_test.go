package test

import (
	"math"
	"testing"
	"github.com/arnoldadero/linear-stats/regression"
)

func TestCalculateLinearRegression(t *testing.T) {
	testCases := []struct {
		name              string
		data              []float64
		expectedSlope     float64
		expectedIntercept float64
	}{
		{
			name:              "Positive correlation",
			data:              []float64{1, 2, 3, 4, 5},
			expectedSlope:     1.0,
			expectedIntercept: 1.0,
		},
		{
			name:              "Negative correlation",
			data:              []float64{5, 4, 3, 2, 1},
			expectedSlope:     -1.0,
			expectedIntercept: 5.0,
		},
		{
			name:              "No correlation",
			data:              []float64{1, 1, 1, 1, 1},
			expectedSlope:     0.0,
			expectedIntercept: 1.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slope, intercept := regression.CalculateLinearRegression(tc.data)
			if math.Abs(slope-tc.expectedSlope) > 1e-6 || math.Abs(intercept-tc.expectedIntercept) > 1e-6 {
				t.Errorf("Expected slope: %f, intercept: %f; got slope: %f, intercept: %f",
					tc.expectedSlope, tc.expectedIntercept, slope, intercept)
			}
		})
	}
}
