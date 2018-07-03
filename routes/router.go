package routes

import (
	"github.com/gorilla/mux"
	"github.com/ismayilmalik/golang-blog/controllers"
)

func Routers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeGET).Methods("GET")

	r.HandleFunc("/register", controllers.RegisterGET).Methods("GET")
	r.HandleFunc("/register", controllers.RegisterPOST).Methods("POST")

	r.HandleFunc("/login", controllers.LoginGET).Methods("GET")
	r.HandleFunc("/login", controllers.LoginPOST).Methods("POST")

	return r
}
