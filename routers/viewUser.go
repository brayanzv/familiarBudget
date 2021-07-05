package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/brayanzv/FamiliarBudget2/bd"
)

func ViewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	auxID := vars["user_id"]

	if auxID == "" {
		http.Error(w, "no ha obtenido aux = " + auxID, 0)
		return
	}

	user, err := bd.SearchUser(auxID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el Usuario "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "aplication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
