package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("serveur lancé sur le port 80 !")
	fmt.Println("http://localhost:80")
	http.ListenAndServe(":80", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Récupérer le formulaire envoyé par le client et l'afficher
	t, _ := template.ParseFiles("index.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}
	//On récupère les données du formulaire
	r.ParseForm()
	//On récupère les valeurs du formulaire
	Form_input := r.FormValue("input")
	Form_select := r.FormValue("select")
	Form_textarea := r.FormValue("textarea")
	Form_radio := r.FormValue("radio")
	Form_checkbox := r.PostForm["checkbox"]
	//On les affiche
	fmt.Println("input:", Form_input)
	fmt.Println("select:", Form_select)
	fmt.Println("textarea:", Form_textarea)
	fmt.Println("radio:", Form_radio)
	fmt.Println("checkbox:", Form_checkbox)
	t.Execute(w, nil)
}
