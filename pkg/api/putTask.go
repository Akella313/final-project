package api

import (
	"encoding/json"
	"final-project/pkg/db"
	"net/http"
	"strings"
)

func putTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	task.ID = strings.TrimSpace(task.ID)
	if task.ID == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	task.Title = strings.TrimSpace(task.Title)
	if task.Title == "" {
		writeJSON(w, map[string]string{
			"error": "не указан заголовок задачи",
		})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}
