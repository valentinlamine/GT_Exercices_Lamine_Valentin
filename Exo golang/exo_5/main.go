package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/2", Handler_2)
	http.HandleFunc("/3", Handler_3)
	http.HandleFunc("/4", Handler_4)
	fmt.Println("serveur lancé sur le port 80 !")
	fmt.Println("http://localhost:80")
	http.ListenAndServe(":80", Custom404Handler(nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func Handler_2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "2.html")
}

func Handler_3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "3.html")
}

func Handler_4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "4.html")
}

func Custom404Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h != nil {
			h.ServeHTTP(w, r)
		}
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
		} else {
			http.ServeFile(w, r, "404.html")
		}
	})
}
