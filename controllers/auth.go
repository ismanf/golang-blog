package controllers

import (
	"time"
	"github.com/ismayilmalik/golang-blog/infrastructure/session"
	"fmt"
	"net/http"

	"github.com/ismayilmalik/golang-blog/models"

	"github.com/ismayilmalik/golang-blog/infrastructure/veng"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {
	veng.Render(w, "login", nil)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	password := r.FormValue("password")
	author := &models.Author{}
	author.Login = r.FormValue("login")
	err := author.GetByLogin()

	if err != nil {
		fmt.Println(err)
	}

	if author.Password != password {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	token, _ := session.New(author.Login)
	fmt.Println("token", token)
	http.SetCookie(w, &http.Cookie {
		Name: "session_token",
		Value: token,
		Expires: time.Now().Add(10 * time.Minute),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RegisterGET(w http.ResponseWriter, r *http.Request) {

	veng.Render(w, "register", nil)
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	author := &models.Author{}
	author.FirstName = r.FormValue("firstName")
	author.LastName = r.FormValue("lastName")
	author.Login = r.FormValue("login")
	author.Password = r.FormValue("password")
	author.Email = r.FormValue("email")

	err := author.Create()

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
