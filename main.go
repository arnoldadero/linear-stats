package main

import (
	"fmt"
	"os"

	"github.com/arnoldadero/linear-stats/regression"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as a command-line argument")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := regression.ReadDataFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(data) < 2 {
		fmt.Println("Not enough data points to calculate regression and correlation")
		os.Exit(1)
	}

	slope, intercept := regression.CalculateLinearRegression(data)
	correlation := regression.CalculatePearsonCorrelation(data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlation)
}
