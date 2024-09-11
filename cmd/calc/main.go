package main
import "/sujet/calc"

import (
	"bufio"
	"fmt"
	"os"
	"calc" // Importe le package "calc" qui contient les autres fonctionnalités
)

func main() {
	// Initialiser la mémoire
	memory := calc.NewMemory()

	// Scanner pour recevoir les entrées de l'utilisateur
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenue dans le calculateur. Tapez 'exit' pour quitter.")

	for {
		fmt.Print("calc> ")
		scanner.Scan()
		input := scanner.Text()

		// Condition de sortie
		if input == "exit" {
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
