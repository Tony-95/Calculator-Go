package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Tony-95/Calculator-Go/calc"
)

func main() {
	// Initialiser la mémoire
	memory := calc.NewMemory()

	// Scanner pour recevoir les entrées de l'utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenue dans le calculateur. Tapez 'quitter' pour quitter.")

	for {
		fmt.Print("calc> ")
		scanner.Scan()
		input := scanner.Text()

		// Condition de sortie
		if input == "quitter" {
			break
		}

		// Essayer d'effectuer un calcul basique
		result, err := calc.BasicCalculate(input)
		if err != nil {
			// Si pas de résultat basique, tenter exponentiation
			result, err = calc.PowerCalculate(input)
		}

		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Println("Résultat:", result)
			// Stocker le résultat en mémoire sous une clé spécifique (par exemple, calc1, calc2, etc.)
			key := fmt.Sprintf("calc%d", len(memory.Store)+1)
			memory.StoreResult(key, result)
		}
	}

	fmt.Println("Merci d'avoir utilisé le calculateur!")
}
