package handlers

import (
	"log"
	"net/http"
	"os"

	middlew "github.com/brayanzv/FamiliarBudget2/middlew"
	routers "github.com/brayanzv/FamiliarBudget2/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handler() {

	router := mux.NewRouter()

	router.HandleFunc("/login", routers.Login).Methods("POST")
	router.HandleFunc("/register",routers.Register).Methods("POST")

	//super users
	router.HandleFunc("/users/{user_id}", middlew.ValidatorJWT(routers.ViewUser)).Methods("GET") //ve la informacion de un usuario especifico
	router.HandleFunc("/users/{user_id}", middlew.ValidatorJWT(routers.ModifyUser)).Methods("PUT") //actualiza un usario especifico
	router.HandleFunc("/users", middlew.ValidatorJWT(routers.UsersList)).Methods("GET") //get all users
	router.HandleFunc("/users/{user_id}", middlew.ValidatorJWT(routers.DeleteUser)).Methods("DELETE")

	//users
	router.HandleFunc("/family/{codFamily}/users", middlew.ValidatorJWT(routers.UsersListCodFamily)).Methods("GET")
	router.HandleFunc("/family/{codFamily}/users/{user_id}", middlew.ValidatorJWT(routers.UsersFamilyModify)).Methods("PUT")

	//user detalle
	router.HandleFunc("/user/{user_id}/detail",middlew.ValidatorJWT(routers.DetailRegister)).Methods("POST") //registrar una transaccion
	router.HandleFunc("/user/{user_id}/detail/{codFamily}", middlew.ValidatorJWT(routers.GetDetailsUser)).Methods("GET") // obtiene todos los detalles del usuario especifico
	router.HandleFunc("/user/detail/{codFamily}", middlew.ValidatorJWT(routers.GetDetailsFamily)).Methods("GET")
	router.HandleFunc("/user/{user_id}/detail/{codFamily}", middlew.ValidatorJWT(routers.DeleteDetailUSer)).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
