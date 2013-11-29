
package engine

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type webhook struct {
	After string `json:"after"`
	HeadCommit headCommit `json:"head_commit"`
	Repository repo `json:"repository"`
}

type headCommit struct {
	Added []string `json:"added"`
	Modified []string `json:"modified"`
	Removed []string `json:"removed"`
	Author author `json:"author"`
	Committer committer `json:"committer"`
	Message string `json:"message"`
}

type author struct {
	Email string `json:"email"`
	Name string `json:"name"`
	UserName string `json:"username"`
}

type committer struct {
	Email string `json:"email"`
	Name string `json:"name"`
	UserName string `json:"username"`
}

type repo struct {
	CreatedAt int64 `json:"created_at"`
	Name string `json:"name"`
	Description string `json:"description"`
	MasterBranch string `json:"master_branch"`
	Url string `json:"url"`
}

type Configuration struct {
	port string
	logFile string
}

func Start(c Configuration) {
	log.Printf("Starting Deamon on Port: %s", c.port)

	mux := mux.NewRouter()
	mux.HandleFunc("/bounce/{name}/{version}", http.HandlerFunc(bounce))
	mux.HandleFunc("/hook", http.HandlerFunc(hook))
	http.Handle("/", mux)

	log.Println("Listening...")
	http.ListenAndServe(c.port, nil)
}

func bounce(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Home")
	params := mux.Vars(r)
	name := params["name"]
	version := params["version"]
	w.Write([]byte("Bouncing " + name + "-" + version))
}

func hook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t webhook

	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	} else {
		log.Printf("Configuration change for %s was pushed to Github. Bounce testing application now", t.Repository.Name)
	}
}

