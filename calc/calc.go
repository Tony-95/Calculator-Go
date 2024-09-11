package calc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Fonction principale de calcul
func Calculate(input string) (float64, error) {
	input = strings.TrimSpace(input)

	// Addition
	if strings.Contains(input, "+") {
		parts := strings.Split(input, "+")
		if len(parts) == 2 {
			left, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			right, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if err1 == nil && err2 == nil {
				return left + right, nil
			}
		}
	}

	// Multiplication
	if strings.Contains(input, "*") {
		parts := strings.Split(input, "*")
		if len(parts) == 2 {
			left, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			right, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if err1 == nil && err2 == nil {
				return left * right, nil
			}
		}
	}

	// Division
	if strings.Contains(input, "/") {
		parts := strings.Split(input, "/")
		if len(parts) == 2 {
			left, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			right, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if err1 == nil && err2 == nil {
				if right == 0 {
					return 0, fmt.Errorf("division par zéro")
				}
				return left / right, nil
			}
		}
	}

	// Exponentiation
	if strings.Contains(input, "^") {
		parts := strings.Split(input, "^")
		if len(parts) == 2 {
			base, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			exponent, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if err1 == nil && err2 == nil {
				return math.Pow(base, exponent), nil
			}
		}
	}

	return 0, fmt.Errorf("opération non reconnue")
}
