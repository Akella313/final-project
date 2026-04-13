package server

import (
	"final-project/pkg/api"
	"net/http"
	"os"
)

func Run() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	api.Init()
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":"+port, nil)
}
