package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL,
    title VARCHAR NOT NULL,
    comment TEXT,
    repeat VARCHAR
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	install := false
	if err != nil {
		if os.IsNotExist(err) {
			install = true
		} else {
			return err
		}
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			db.Close()
			return err
		}
	}

	DB = db
	return nil
}
