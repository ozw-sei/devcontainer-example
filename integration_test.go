package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSetupHandler(t *testing.T) {
	dsn := "project:password@tcp(mysql:3306)/project?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	setupHandler(db)

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.DefaultServeMux
	handler.ServeHTTP(rr, req)

	if http.StatusOK != rr.Code {
		t.Errorf("rr.Code is not 200")
	}

	var tasks []models.Task
	err = json.NewDecoder(rr.Body).Decode(&tasks)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 10 {
		t.Errorf("len(tasks) is not 10")
	}
}
