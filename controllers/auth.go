package controllers

import (
	"net/http"

	"github.com/ismayilmalik/golang-blog/infrastructure/veng"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {

	veng.Render(w, "login", nil)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {

	veng.Render(w, "login", nil)
}

func RegisterGET(w http.ResponseWriter, r *http.Request) {

	veng.Render(w, "register", nil)
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {

	veng.Render(w, "register", nil)
}
