# Linear Regression and Correlation Calculator

This Go program calculates and prints the Linear Regression Line and Pearson Correlation Coefficient from numeric data in a file.

## Usage

```
go run main.go data.txt
```

The data file should contain one number per line.

## Output

The program will output the Linear Regression Line equation and the Pearson Correlation Coefficient in the following format:

```
Linear Regression Line: y = <slope>x + <intercept>
Pearson Correlation Coefficient: <coefficient>
```

- Linear Regression Line values (slope and intercept) are displayed with 6 decimal places.
- Pearson Correlation Coefficient is displayed with 10 decimal places.

## Error Handling

The program includes error handling for:
- Missing file path argument
- File reading errors
- Invalid numeric data in the file
- Insufficient data points (less than 2)

## Implementation Details

The program uses the following formulas for calculations:

1. Linear Regression Line (y = mx + b):
   - Slope (m) = (n∑xy - ∑x∑y) / (n∑x² - (∑x)²)
   - Intercept (b) = (∑y - m∑x) / n

2. Pearson Correlation Coefficient (r):
   r = (n∑xy - ∑x∑y) / sqrt((n∑x² - (∑x)²) * (n∑y² - (∑y)²))

Where:
- n is the number of data points
- x is the index of each data point (0, 1, 2, ...)
- y is the value of each data point

## License

This project is open-source and available under the MIT License.