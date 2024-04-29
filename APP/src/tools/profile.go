package tools

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetIdByCookie(w http.ResponseWriter, r *http.Request) (int, error) {
	CheckCookie(w, r)
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusInternalServerError)
		return 0, err
	}
	idData, err := http.Get("http://localhost:4000/get-id-by-cookie?cookie=" + cookie.Value)
	if err != nil {
		// Gérer l'erreur, par exemple en affichant un message d'erreur sur la page ou en redirigeant l'utilisateur
		http.Error(w, "Erreur lors de la récupération des informations utilisateur", http.StatusInternalServerError)
		return 0, err
	}
	defer idData.Body.Close()

	idBytes, err := ioutil.ReadAll(idData.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusInternalServerError)
		return 0, err
	}

	id, err := strconv.Atoi(string(idBytes))
	if err != nil {
		http.Error(w, "Erreur lors de la conversion de l'ID en entier"+string(idBytes), http.StatusInternalServerError)
		return 0, err
	}
	return id, nil
}
