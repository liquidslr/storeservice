package routes

import (
	"log"
	"net/http"
)

// Server to start an HTTP listener at mentioned port)
func Server(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("[ERROR] Unable to start HTTP listener at port " + port)
		log.Println("[ERROR] Error: " + err.Error())
	}
}
