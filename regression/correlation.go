package regression

import "math"

// CalculatePearsonCorrelation calculates the Pearson correlation coefficient
func CalculatePearsonCorrelation(data []float64) float64 {
	n := float64(len(data))
	sumX, sumY, sumXY, sumXX, sumYY := 0.0, 0.0, 0.0, 0.0, 0.0

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
		sumYY += y * y
	}

	// Calculate Pearson correlation coefficient
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumXX - sumX*sumX) * (n*sumYY - sumY*sumY))

	return numerator / denominator
}
