package main

import (
	"btpn-ta/config"
	homecontroller "btpn-ta/controllers/homeController"
	logincontroller "btpn-ta/controllers/loginController"
	registercontroller "btpn-ta/controllers/registerController"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/home", homecontroller.Index)

	http.HandleFunc("/login", logincontroller.Index)
	http.HandleFunc("/regis", registercontroller.Index)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
