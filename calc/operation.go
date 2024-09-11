package calc

import (
	"math"
	"strconv"
	"strings"
	"fmt"
)

// Fonction pour effectuer l'exponentiation
func PowerCalculate(input string) (float64, error) {
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
