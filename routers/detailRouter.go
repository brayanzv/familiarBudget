package routers

import (
	"encoding/json"
	"github.com/brayanzv/FamiliarBudget2/bd"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/brayanzv/FamiliarBudget2/models"
)

//valida la informacion enviada en el body y sus credenciales
func DetailRegister(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	idUs:=vars["user_id"]
	idU,_:=strconv.Atoi(idUs)

	if uint(idU)!=IdUsers{
		http.Error(w, "id invalido",403)
		return
	}

	var t models.Details

	err := json.NewDecoder(r.Body).Decode(&t)
	t.Id_user=idU
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	status, err :=bd.InsertDetail(t)
	if status==false{
		return
		http.Error(w,"No se pudo registrar el Detalle ", 400)
	}
	//este va a ser otro comment
	w.WriteHeader(http.StatusCreated)
}
func GetDetailsUser(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	IdUs:=vars["user_id"]
	codFamily := vars["codFamily"]
	idU,_ := strconv.Atoi(IdUs)

	if codFamily != CodFamiliar{
		http.Error(w,"Error en credenciales", 403)
		return
	}
	result, status := bd.GetDetailsUserDB(uint(idU),codFamily)

	if status == false{
		http.Error(w, "No se pudo realizar la busqueda", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func GetDetailsFamily(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	codFamily := vars["codFamily"]

	if codFamily != CodFamiliar{
		http.Error(w,"Credenciales incorrectas",403)
		return
	}
	result, status := bd.GetDetailsFamilyDB(codFamily)

	if status == false{
		http.Error(w, "No se pudo realizar la busqueda", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
