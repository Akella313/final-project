package api

import (
	"encoding/json"
	"final-project/pkg/db"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func addTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	task.Title = strings.TrimSpace(task.Title)
	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "не указан заголовок задачи"})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{
		"id": strconv.FormatInt(id, 10),
	})
}

func checkDate(task *db.Task) error {
	now := time.Now()

	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if task.Date == "" {
		task.Date = today.Format("20060102")
	}
	taskDate, err := time.ParseInLocation("20060102", task.Date, today.Location())
	if err != nil {
		return err
	}

	var next string
	if task.Repeat != "" {
		next, err = NextDate(today, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if taskDate.Before(today) {
		if task.Repeat == "" {
			task.Date = today.Format("20060102")
		} else {
			task.Date = next
		}
	}

	return nil
}
