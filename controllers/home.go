package controllers

import (
	"fmt"
	"github.com/ismayilmalik/golang-blog/infrastructure/session"
	"net/http"

	"github.com/ismayilmalik/golang-blog/infrastructure/veng"
	"github.com/ismayilmalik/golang-blog/models"
)

func HomeGET(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Value
	fmt.Println("sessionToken", sessionToken)
	isValid, err := session.IsValid(sessionToken) 

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	article := &models.Article{}
	articles, _ := article.GetAll()
	veng.Render(w, "home", articles)
}
