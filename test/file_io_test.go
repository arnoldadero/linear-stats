package test

import (
	"os"
	"testing"
	"github.com/arnoldadero/linear-stats/regression"
)

func TestReadDataFromFile(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "test_data_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test data to the file
	testData := []string{"1.5", "2.3", "3.7", "4.2", "5.1"}
	for _, data := range testData {
		if _, err := tmpfile.WriteString(data + "\n"); err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
	}
	tmpfile.Close()

	// Test reading data from the file
	data, err := regression.ReadDataFromFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadDataFromFile failed: %v", err)
	}

	// Check if the data matches the expected values
	expected := []float64{1.5, 2.3, 3.7, 4.2, 5.1}
	if len(data) != len(expected) {
		t.Errorf("Expected %d data points, got %d", len(expected), len(data))
	}
	for i, value := range data {
		if value != expected[i] {
			t.Errorf("Data point %d: expected %f, got %f", i, expected[i], value)
		}
	}
}

func TestReadDataFromFileErrors(t *testing.T) {
	// Test non-existent file
	_, err := regression.ReadDataFromFile("non_existent_file.txt")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}

	// Test file with invalid data
	tmpfile, err := os.CreateTemp("", "invalid_data_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString("1.5\n2.3\ninvalid\n4.2\n"); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpfile.Close()

	_, err = regression.ReadDataFromFile(tmpfile.Name())
	if err == nil {
		t.Error("Expected error for invalid data, got nil")
	}
}
