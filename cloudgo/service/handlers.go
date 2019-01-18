package service

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/unrolled/render"
)

func apiTimeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		whatTime := time.Now().String()
		formatter.JSON(w, http.StatusOK, struct {
			Randnum string
		}{Randnum: whatTime})
	}
}

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		template := template.Must(template.New("index.html").ParseFiles("./templates/index.html"))
		randnum := rand.Intn(10)
		_ = template.Execute(w, struct {
			MyName  string
			Randnum int
		}{MyName: "stevewuzhi", Randnum: randnum})
	}
}

func checkform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username"))
	password := template.HTMLEscapeString(r.Form.Get("password"))
	randnum := rand.Intn(10)
	t := template.Must(template.New("signin.html").ParseFiles("./templates/signin.html"))
	err := t.Execute(w, struct {
		Username string
		Password string
		Randnum  int
	}{Username: username, Password: password, Randnum: randnum})
	if err != nil {
		panic(err)
	}
}

func inDevelopmentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("This is in development."))
}
