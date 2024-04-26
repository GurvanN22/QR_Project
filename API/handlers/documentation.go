package handlers

import (
	"net/http"
)

func Documentation(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/swagger.json")
}
