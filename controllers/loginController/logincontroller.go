package logincontroller

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("views/login/login.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
