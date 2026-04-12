package server

import (
	"final-project/pkg/api"
	"net/http"
	"os"
)

// Функция Run() запускает сервер с портом из переменной окружения TODO_PORT
// Если переменная не задана, используется порт по умолчанию 7540
func Run() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	api.Init()
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":"+port, nil)
}
