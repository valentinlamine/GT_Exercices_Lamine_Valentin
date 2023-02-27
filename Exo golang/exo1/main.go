package main

import (
	"fmt"
	"strconv"
)

func main() {
	JeuxAmulettes()
}

func JeuxAmulettes() {
	var TourJoueur int
	fmt.Println("Combien d'amulettes voulez-vous ? (entre 4 et 150)")
	var nbAllumettes = Input(4, 150)
	fmt.Println("Combien de joueurs voulez-vous ? (entre 2 et 5)")
	var nbJoueurs = Input(2, 5)

	for nbAllumettes > 0 {
		TourJoueur = TourJoueur % nbJoueurs
		fmt.Println("C'est au tour du joueur " + strconv.Itoa(TourJoueur+1))
		fmt.Println("Il reste", nbAllumettes, "amulettes")
		fmt.Println("Combien d'amulettes voulez-vous enlever ? (entre 1 et 3)")
		nbAllumettes -= Input(1, 3)
		if nbAllumettes <= 0 {
			fmt.Println("Le joueur " + strconv.Itoa(TourJoueur+1) + " a perdu !")
		}
		TourJoueur++
	}
}

func Input(nb_min, nb_max int) int {
	var choix string
	fmt.Scan(&choix)                    // On récupère la saisie de l'utilisateur
	chiffre, err := strconv.Atoi(choix) // On convertit la saisie en entier
	for err != nil || chiffre < nb_min || chiffre > nb_max {
		fmt.Println("Veuillez entrer un nombre entre", nb_min, "et", nb_max)
		fmt.Scan(&choix)
		chiffre, err = strconv.Atoi(choix)
	}
	return chiffre
}
