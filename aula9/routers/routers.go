package routers

import (
	personcontroller "../controllers"
	"github.com/gorilla/mux"
)

// GetRouters pega todas as rotas
func GetRouters() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/contato", personcontroller.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", personcontroller.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", personcontroller.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", personcontroller.DeletePerson).Methods("DELETE")

	return router
}
