package calc

import (
	"strconv"
	"strings"
	"fmt"
)

// Fonction pour effectuer un calcul basique
func BasicCalculate(input string) (float64, error) {
	input = strings.TrimSpace(input)

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

	return 0, fmt.Errorf("opération non reconnue")
}
