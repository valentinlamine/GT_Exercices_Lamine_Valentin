package main

/*
Créer un tableau à n dimensions,
n étant défini de manièrealéatoire (minimum 3),
remplie de valeurs aléatoires depuis lego puis l'afficher en web sous format de tableau.
*/

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var tableau [][]int

func main() {
	rand.Seed(time.Now().UnixNano())
	var taille = rand.Intn(10) + 3
	fmt.Println("Taille du tableau :", taille)
	for i := 0; i < taille; i++ {
		tableau = append(tableau, []int{})
		for j := 0; j < taille; j++ {
			tableau[i] = append(tableau[i], rand.Intn(100))
		}
	}
	http.HandleFunc("/", IndexHandler)
	fmt.Println("serveur lancé sur le port 80 !")
	fmt.Println("http://localhost:80")
	http.ListenAndServe(":80", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	//Ajouter dans la balise <body> le tableau en html
	t.Execute(w, tableau) //On envoie le tableau en paramètre
}
