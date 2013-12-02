
package engine

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type instance struct {
	Name string
	Type string
	Version string
}

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
	Port string `json:"port"`
	LogFile string `json:"logFile"`
}

func Start(c* Configuration) {
	log.Printf("Starting Deamon on Port: %s", c.Port)

	mux := mux.NewRouter()
	mux.HandleFunc("/bounce/{type}/{name}/{version}", http.HandlerFunc(bounce))
	mux.HandleFunc("/hook", http.HandlerFunc(hook))
	http.Handle("/", mux)

	log.Println("Listening...")
	http.ListenAndServe(":" + c.Port, nil)
}

func bounce(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Home")
	params := mux.Vars(r)
	var i instance

	i.Name = params["name"]
	i.Version = params["version"]
	i.Type = params["type"]

	w.Write([]byte("Bouncing " + i.Name + "Version: " + i.Type + " " + i.Type ))
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

