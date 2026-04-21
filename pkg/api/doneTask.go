package api

import (
	"final-project/pkg/db"
	"net/http"
	"strings"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if strings.TrimSpace(task.Repeat) == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}
	} else {
		next, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := db.UpdateDate(next, id); err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}
	}

	writeJSON(w, map[string]string{})
}
