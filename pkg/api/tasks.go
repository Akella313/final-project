package api

import (
	"final-project/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": "метод не поддерживается",
		})
		return
	}
	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
