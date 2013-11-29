
package main

import (
	"log"
	"encoding/json"
	"os"
	"github.com/samilton/bouncer/engine"
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


func main() {
	config, err := os.Open("conf.json")

	log.Println("Starting Web Bouncer Daemon")
	if err == nil {
		decoder := json.NewDecoder(config)
		configuration := &Configuration{}
		decoder.Decode(&configuration)
		engine.Start(&configuration)
	} else {
		panic(err)
	}
}


