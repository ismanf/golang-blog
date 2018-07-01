package main

import (
	"log"
	"github.com/ismayilmalik/golang-blog/veng"
	"net/http"
)

func main() {
	veng.Initialize(nil)
	http.HandleFunc("/home", ServeHome)
	http.HandleFunc("/article", ServeArticle)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	veng.Render(w, "home", "")
}

func ServeArticle(w http.ResponseWriter, r *http.Request) {
	veng.Render(w, "article", "")
}

