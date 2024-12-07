package models_test

import (
	"project/models"
	"testing"
)

func Test_Tasks(t *testing.T) {
	task := models.Task{}

	if "Hello" != task.Hello() {
		t.Errorf("task.Hello() is failedww")
	}
}
