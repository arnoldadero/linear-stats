package regression

// CalculateLinearRegression calculates the slope and intercept of the linear regression line
func CalculateLinearRegression(data []float64) (slope, intercept float64) {
	n := float64(len(data))
	sumX, sumY, sumXY, sumXX := 0.0, 0.0, 0.0, 0.0

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
	}

	// Calculate slope (m) and intercept (b) for y = mx + b
	slope = (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	intercept = (sumY - slope*sumX) / n

	return slope, intercept
}
