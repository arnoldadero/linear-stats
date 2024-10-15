package regression

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadDataFromFile reads numeric data from a file and returns a slice of float64 values
func ReadDataFromFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number in file: %v", err)
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
