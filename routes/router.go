package routes

import (
	"github.com/gorilla/mux"
	"github.com/ismayilmalik/golang-blog/controllers"
)

func Routers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.RegisterGET).Methods("GET")

	return r
}