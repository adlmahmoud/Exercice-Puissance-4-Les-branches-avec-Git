package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var grille [6][7]string
	joueur := "X"
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			grille[i][j] = "."
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🎮 PUISSANCE 4")
	fmt.Println("Joueur 1: X, Joueur 2: O")
	for {
		afficherGrille(grille)
		fmt.Printf("Tour du joueur %s\n", joueur)
		colonne := demanderColonne(scanner)
		if colonne == -1 {
			fmt.Println("Erreur de saisie ! Réessaye.")
			continue
		}

		if !placerPion(&grille, colonne, joueur) {
			fmt.Println("Colonne pleine ! Réessaye.")
			continue
		}

		if verifierVictoire(grille, joueur) {
			afficherGrille(grille)
			fmt.Printf("🎉 Joueur %s gagne ! 🎉\n", joueur)
			break
		}
		if grillePleine(grille) {
			afficherGrille(grille)
			fmt.Println("🤝 Match nul ! 🤝")
			break
		}
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
	}
}

func afficherGrille(g [6][7]string) {
	fmt.Println("\n  1 2 3 4 5 6 7")
	for i := 0; i < 6; i++ {
		fmt.Print("| ")
		for j := 0; j < 7; j++ {
			fmt.Print(g[i][j] + " ")
		}
		fmt.Println("|")
	}
	fmt.Println("-----------------")
}

func demanderColonne(scanner *bufio.Scanner) int {
	fmt.Print("Choisis une colonne (1-7): ")
	if scanner.Scan() {
		texte := strings.TrimSpace(scanner.Text())
		colonne, err := strconv.Atoi(texte)
		if err != nil {
			return -1
		}

		if colonne >= 1 && colonne <= 7 {
			return colonne - 1
		}
	}
	return -1
}

func placerPion(g *[6][7]string, col int, joueur string) bool {
	for ligne := 5; ligne >= 0; ligne-- {
		if g[ligne][col] == "." {
			g[ligne][col] = joueur
			return true
		}
	}
	return false
}

func verifierVictoire(g [6][7]string, joueur string) bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			if g[i][j] == joueur && g[i][j+1] == joueur &&
				g[i][j+2] == joueur && g[i][j+3] == joueur {
				return true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 7; j++ {
			if g[i][j] == joueur && g[i+1][j] == joueur &&
				g[i+2][j] == joueur && g[i+3][j] == joueur {
				return true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if g[i][j] == joueur && g[i+1][j+1] == joueur &&
				g[i+2][j+2] == joueur && g[i+3][j+3] == joueur {
				return true
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 3; j < 7; j++ {
			if g[i][j] == joueur && g[i+1][j-1] == joueur &&
				g[i+2][j-2] == joueur && g[i+3][j-3] == joueur {
				return true
			}
		}
	}

	return false
}
func grillePleine(g [6][7]string) bool {
	for j := 0; j < 7; j++ {
		if g[0][j] == "." {
			return false
		}
	}
	return true
}
