package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func main() {
	hub := newHub()
	go hub.run()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods(http.MethodGet)
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
