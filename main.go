package main

import (
	"final-project/pkg/db"
	"final-project/pkg/server"
)

func main() {
	// Инициализация базы данных
	err := db.Init("./scheduler.db")
	if err != nil {
		panic(err)
	}

	// Запускает сервер
	server.Run()
}
