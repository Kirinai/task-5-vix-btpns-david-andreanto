package registercontroller

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("views/regis/regis.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
