package handlers

import (
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

func (t *Tasks) CreateTask(rw http.ResponseWriter, r *http.Request) {

	task := r.Context().Value(TaskKey{}).(data.Task)
	t.db.Create(task)
}
