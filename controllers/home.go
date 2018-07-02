package controllers

import (
	"github.com/ismayilmalik/golang-blog/infrastructure/veng"
	"net/http"
)

func HomeGET(w http.ResponseWriter, r *http.Request) {

	veng.Render(w, "home", "")
}