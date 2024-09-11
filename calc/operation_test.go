package calc

import (
	"fmt"
	"testing"
)

// Fonction utilitaire pour la multiplication
func Multiply(a, b float64) float64 {
	return a * b
}

// Fonction utilitaire pour la division
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division par zéro")
	}
	return a / b, nil
}

// Tests unitaires pour la multiplication
func Test_Multiply(t *testing.T) {
	result := Multiply(5, 3)
	expected := 15.0
	if result != expected {
		t.Errorf("Échec du test Multiply, attendu: %f, obtenu: %f", expected, result)
	}
}

// Tests unitaires pour la division
func Test_Divide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Erreur inattendue lors de la division: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Échec du test Divide, attendu: %f, obtenu: %f", expected, result)
	}

	// Test pour la division par zéro
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("La division par zéro aurait dû échouer, mais elle a réussi")
	}
}
