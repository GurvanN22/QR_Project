package handlers

import (
	"database/sql"
	"log"
	"net/http"
)

func SaveCookie(w http.ResponseWriter, r *http.Request) {
	// Vérifier la méthode de requête
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer le cookie de session de la requête HTTP
	sessionID := r.FormValue("session_id")
	if sessionID == "" {
		http.Error(w, "Session ID manquant", http.StatusBadRequest)
		return
	}

	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", "db/data.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Enregistrer le cookie de session dans la base de données
	_, err = db.Exec("INSERT INTO session_cookie (session_id) VALUES (?)", sessionID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de l'enregistrement du cookie dans la base de données", http.StatusInternalServerError)
		return
	}

	// Répondre avec un message de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cookie enregistré avec succès dans la base de données."))
}
