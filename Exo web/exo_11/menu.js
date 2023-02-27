function SetWeight() {
    //récupérer la taille de l'élément dropdown
    var weight = document.querySelector(".dropdown").offsetWidth;
    //définir la taille de l'élément dropdown menu
    document.querySelectorAll(".dropdown-menu").forEach(function (element) {
        element.style.width = weight + "px";
    });
}

//charge la fonction SetWeight() lorsque la page est chargée
window.onload = SetWeight;
//charge la fonction SetWeight() lorsque la taille de la fenêtre est modifiée
window.onresize = SetWeight;