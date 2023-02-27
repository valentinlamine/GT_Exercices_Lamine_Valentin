package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	boucle := true
	fmt.Println("Pour un fonctionnemment idéal vérifier que votre antivirus n'interfère pas avec le programme")
	for boucle {
		var choix int
		fmt.Println("\n\n\nBienvenue dans le menu de gestion de fichiers texte !")
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 - Récupérer le texte contenu dans un fichier")
		fmt.Println("2 - Ajouter du texte dans le fichier")
		fmt.Println("3 - Supprimer tout le contenu du fichier")
		fmt.Println("4 - Remplacer le contenu du fichier par le texte donné")
		fmt.Println("5 - Quitter le programme")
		choix = Input(1, 5)
		switch choix {
		case 1:
			LireContenuFichier()
		case 2:
			AjouterContenuFichier()
		case 3:
			ViderContenuFichier()
		case 4:
			RemplacerContenuFichier()
		case 5:
			fmt.Println("Au revoir !")
			boucle = false
		}
	}
}

func LireContenuFichier() {
	fichier := OuvrirFichier()
	defer fichier.Close()
	fmt.Print("\n\n\nVoici le contenu du fichier :\n\n")
	contenu := make([]byte, 1024)
	var buffer bytes.Buffer
	for {
		n, err := fichier.Read(contenu)
		if err == io.EOF {
			buffer.Write(contenu[:n])
			break
		}
		if err != nil {
			fmt.Println("Une erreur est survenue lors de la lecture du fichier")
			fmt.Println(err)
			break
		}
		buffer.Write(contenu[:n])
	}
	fmt.Println(buffer.String())
}

func AjouterContenuFichier() {
	fichier := OuvrirFichier()
	defer fichier.Close()
	fichier.Seek(0, 2)
	fichier.WriteString("\n")
	fmt.Println("\n\n\nVeuillez entrer le texte à ajouter au fichier :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texte := scanner.Text()
		_, err := fichier.WriteString(texte)
		if err != nil {
			fmt.Println("Une erreur est survenue lors de l'écriture dans le fichier")
			fmt.Println(err)
		}
	}
}

func ViderContenuFichier() {
	fichier := OuvrirFichier()
	fmt.Println("Etes-vous sûr de vouloir vider le fichier ? (Oui-1/Non-2)")
	choix := Input(1, 2)
	if choix == 2 {
		return
	}
	defer fichier.Close()
	fichier.Truncate(0)
	fmt.Println("Le fichier a été vidé")
}

func RemplacerContenuFichier() {
	fichier := OuvrirFichier()
	defer fichier.Close()
	fichier.Truncate(0)
	fmt.Println("\n\n\nVeuillez entrer le texte à ajouter au fichier :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texte := scanner.Text()
		_, err := fichier.WriteString(texte)
		if err != nil {
			fmt.Println("Une erreur est survenue lors de l'écriture dans le fichier")
			fmt.Println(err)
		}
	}
}

func OuvrirFichier() *os.File {
	var nomFichier string
	fmt.Println("Veuillez entrer le nom du fichier à ouvrir :")
	fmt.Scan(&nomFichier)
	fichier, err := os.OpenFile(nomFichier, os.O_RDWR, 0755)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			fmt.Println("Le fichier n'existe pas")
		default:
			fmt.Println("Une erreur est survenue lors de l'ouverture du fichier")
			fmt.Println(err)
		}
		return OuvrirFichier()
	}
	return fichier
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
