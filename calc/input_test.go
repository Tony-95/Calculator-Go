package calc

import (
	"strconv"
	"testing"
)

// Fonction utilitaire pour convertir l'entrée en valeur numérique
func ConvertInputToValue(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

// Test unitaire pour la conversion des entrées utilisateur
func Test_ConvertInputToValue(t *testing.T) {
	// Cas de test valide
	val, err := ConvertInputToValue("42.5")
	if err != nil || val != 42.5 {
		t.Errorf("Échec de la conversion de '42.5', attendu: 42.5, obtenu: %f, erreur: %v", val, err)
	}

	// Cas de test pour une entrée invalide
	_, err = ConvertInputToValue("abc")
	if err == nil {
		t.Error("La conversion de 'abc' aurait dû échouer, mais elle a réussi")
	}
}
