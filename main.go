package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/models"

	_ "github.com/go-sql-driver/mysql"
)

func getTasks(db *sql.DB) ([]models.Task, error) {
	// Taskを全件取得する
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		err := rows.Scan(&task.ID, &task.Name, &task.IsCompleted)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func setupHandler(db *sql.DB) {

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := getTasks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(tasks) == 0 {
			for i := 0; i < 10; i++ {
				_, err := db.Exec("INSERT INTO tasks (name) VALUES (?)", fmt.Sprintf("Task %d", i+1))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			tasks, err = getTasks(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

}

func main() {
	// mysql に接続してTaskを10個くらい作成する

	dsn := "project:password@tcp(mysql:3306)/project?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	setupHandler(db)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
