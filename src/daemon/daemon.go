
package daemon

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Configuration struct {
	port string
	logFile string
}

func start(c Configuration) {
	log.Printf("Starting Deamon on Port: %s", c.port)

	mux := mux.NewRouter()
	mux.HandleFunc("/bounce/{name}/{version}", http.HandlerFunc(bounce))
	mux.HandleFunc("/hook", http.HandlerFunc(hook))
	http.Handle("/", mux)

	log.Println("Listening...")
	http.ListenAndServe(c.port, nil)
}

