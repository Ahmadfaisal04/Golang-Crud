package main

import (
	"net/http"
	"pasien/config"
	"pasien/controllers"
)

func main() {
	config.Conn()
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8000", nil)
}
