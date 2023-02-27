const series_object = [
    {
        'name': 'lupin',
        'seasons': 2,
        'episodes': 5,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'dark',
        'seasons': 3,
        'episodes': 8,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'breakingbad',
        'seasons': 5,
        'episodes': 16,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'gameofthrones',
        'seasons': 8,
        'episodes': 10,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'thewalkingdead',
        'seasons': 10,
        'episodes': 16,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'strangerthings',
        'seasons': 3,
        'episodes': 8,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'thesopranos',
        'seasons': 6,
        'episodes': 13,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'thewire',
        'seasons': 5,
        'episodes': 13,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'thegoodplace',
        'seasons': 4,
        'episodes': 13,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'theoffice',
        'seasons': 9,
        'episodes': 24,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'thebigbangtheory',
        'seasons': 12,
        'episodes': 24,
        'duration': 45,
        'ad': 0,
        'total': 0
    },
    {
        'name': 'Not_Found',
        'seasons': 0,
        'episodes': 0,
        'duration': 0,
        'ad': 0,
        'total': 0
    }
]


function PrintSeries(serie) {
    //il y a deux fois la classe .nom_serie dans le html, il faut donc modifier les deux
    document.querySelectorAll(".nom_serie").forEach(function (el) {
        el.innerHTML = serie.name;
    });
    document.querySelector(".nb_saison").innerHTML = serie.seasons;
    document.querySelector(".nb_episode_saison").innerHTML = serie.episodes;
    document.querySelector(".duree_episode").innerHTML = serie.duration;
    document.querySelector(".temps_pub").innerHTML = serie.ad;
    document.querySelector(".total_temps").innerHTML = serie.seasons*serie.episodes*(serie.duration+serie.ad);
}

function Input() {
    //récupérer la valeur de l'input
    var input = document.getElementById("input").value;
    //traier la valeur de l'input pour la mettre en minuscule et enlever les espaces
    var input = input.toLowerCase().replace(/\s/g, '');
    //vérifier si la valeur de l'input correspond à une série
    for (let i = 0; i < series_object.length; i++) {
        if (input == series_object[i].name) {
            PrintSeries(series_object[i]);
            //sortir de la boucle
            break;
        }
        //si la valeur de l'input ne correspond à aucune série, afficher "Série inconnue" dans la console
        else if (i == series_object.length - 1) {
            PrintSeries(series_object[series_object.length - 1]);
        }
    }
}

//charger la fonction myFunction au chargement de la page
window.onload = PrintSeries(series_object[0]);

//charger la fonction Input lorsque l'utilisateur tape entrée dans l'input
document.getElementById("input").addEventListener("keyup", function (event) {
    if (event.keyCode === 13) {
        Input();
    }
});
