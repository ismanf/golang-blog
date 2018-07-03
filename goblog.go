package main

import (
	"github.com/ismayilmalik/golang-blog/infrastructure/cache"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/ismayilmalik/golang-blog/infrastructure/database"
	"github.com/ismayilmalik/golang-blog/infrastructure/veng"
	"github.com/ismayilmalik/golang-blog/routes"
)

type GoBlogApp struct {
	router *mux.Router
}

func (a *GoBlogApp) Run() {
	handler := negroni.Classic()
	handler.UseHandler(a.router)
	log.Fatal(http.ListenAndServe(":3001", handler))
}

func (a *GoBlogApp) Initialize() {

	//Initializing database
	checkError(database.Initialize())

	//Initializing routes
	a.router = routes.Routers()

	//Initialize cache
	cache.Initialize()

	//Initialize view engine
	veng.Initialize(nil)

}

func checkError(err error) {
	if err != nil {
		log.Fatal("GoBlog will exit.Error detected:", err)
	}
}
