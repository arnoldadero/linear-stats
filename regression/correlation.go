package regression

import "math"

// CalculatePearsonCorrelation calculates the Pearson correlation coefficient
func CalculatePearsonCorrelation(data []float64) float64 {
	n := float64(len(data))
	if n < 2 {
		return 0 //not enough data points to calculate correlation coefficient
    }

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
        if denominator == 0 {
            return 0 // division by zero is not allowed
        }
    return numerator / denominator

}
