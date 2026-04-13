package db

import (
	"database/sql"
	"errors"
	"strconv"
)

const (
	queryAddTask    = `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	queryTasks      = `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?`
	queryGetTask    = `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	queryUpdateTask = `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	queryDeleteTask = `DELETE FROM scheduler WHERE id = ?`
	queryUpdateDate = `UPDATE scheduler SET date = ? WHERE id = ?`
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	res, err := DB.Exec(queryAddTask, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return id, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Tasks(limit int) ([]*Task, error) {
	rows, err := DB.Query(queryTasks, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next() {
		var task Task
		var id int

		err := rows.Scan(&id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}

		task.ID = strconv.Itoa(id)
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {

		return nil, err
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	var task Task
	var taskID int

	row := DB.QueryRow(queryGetTask, id)

	err := row.Scan(&taskID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Задача не найдена")
		}
		return nil, err
	}

	task.ID = strconv.Itoa(taskID)
	return &task, nil
}

func UpdateTask(task *Task) error {
	res, err := DB.Exec(queryUpdateTask, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Задача не найдена")
	}

	return nil
}

func DeleteTask(id string) error {
	res, err := DB.Exec(queryDeleteTask, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Задача не найдена")
	}

	return nil
}

func UpdateDate(next string, id string) error {
	res, err := DB.Exec(queryUpdateDate, next, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Задача не найдена")
	}

	return nil
}
