package handlers

import (
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

// swagger:route POST /tasks Tasks createTask
// Create a new TODO task
//
// responses:
//
//	200: productResponse
//	422: errorValidation
//	501: errorResponse
func (t *Tasks) CreateTask(rw http.ResponseWriter, r *http.Request) {

	task := r.Context().Value(KeyProduct{}).(*data.Task)
	t.db.Create(task)
}
