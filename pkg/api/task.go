package api

import (
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTask(w, r)
	case http.MethodGet:
		getTask(w, r)
	case http.MethodPut:
		putTask(w, r)
	case http.MethodDelete:
		deleteTask(w, r)
	default:
		writeJSON(w, map[string]string{
			"error": "метод не поддерживается",
		})
	}
}
