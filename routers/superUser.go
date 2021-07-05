package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/brayanzv/FamiliarBudget2/bd"
	"github.com/brayanzv/FamiliarBudget2/models"
)

//register es la funcion para crear en la bd el registro de usuario
func Register(w http.ResponseWriter, r *http.Request) {
	t:=  models.Users{}

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Login) < 6 {
		http.Error(w, "Error en los datos recibidos, ingrese un login mayor a 5 digitos ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Ingrese una contraseña mayor a 5 digitos ", 400)
		return
	}

	_, found, _ := bd.CheckUser(t.Login)
	if found == true {
		http.Error(w, "Ya existe un usuario registrado con ese login", 400)
		return
	}

	if t.Id_role == 3 {
		cod := bd.CodFamiliar(t.Cod_familiar)
		if cod == false {
			http.Error(w, "Debe ingresar un codigo de familia correcto", 400)
			return
		}
	}

	if t.Id_role == 1 {
		http.Error(w, "Usted no esta autorizado para crear este tipo de usuario", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ModifyUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	auxID := vars["user_id"]
	id,_:= strconv.Atoi(auxID)



		t := models.Users{}
		err := json.NewDecoder(r.Body).Decode(&t)


		if err != nil {
			http.Error(w, "no se pudo decodificar"+err.Error(), 400)
			return
		}

		if Id_roles!=1 {
			http.Error(w, "no se pudo modificar al usuario"+err.Error(), 400)
			return
		}

		var status bool

		status, err = bd.ModifyUsers(t, uint(id))
		if err != nil {
			http.Error(w, "Ocurrio un error al modificar el registro "+err.Error(), 400)
			return
		}
		if status == false {
			http.Error(w, "No se logro modificar el registro vuelva a intentarlo"+err.Error(), 400)
			return
		}
		w.WriteHeader(http.StatusCreated)


}

func UsersList(w http.ResponseWriter, r *http.Request) {

	result, status, c := bd.GetUsers()
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	var search []*models.UserAPI

	for i:=0; i<int(c);i++{

		tempUserApi:= models.UserAPI{}

		tempUserApi.Id = int(result[i].Id)
		tempUserApi.Name=result[i].Name
		tempUserApi.Last_name= result[i].Last_name
		tempUserApi.Id_role = result[i].Id_role
		tempUserApi.Cod_familiar=result[i].Cod_familiar
		tempUserApi.Status=result[i].Status

		search = append(search,&tempUserApi)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(search)
}
func DeleteUser(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	idaux := vars["user_id"]

	if idaux == "" {
		http.Error(w, "no ha obtenido aux = " + idaux, 0)
		return
	}
	id,_:= strconv.Atoi(idaux)
 	err,types :=bd.Delete(id)
	if err!= nil&& types == 2{
		http.Error(w, "no se pudo cerrar la base de datos "+err.Error(), 400)
		return
	}
	if err!= nil && types == 1{
		http.Error(w, "error no se pudo eliminar "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

