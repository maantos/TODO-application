package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maantos/todoApplication/pkg/data"
)

// swagger:route DELETE /tasks/{id} Tasks deleteTask
// Removes task from the database
//
// responses:
//
//	204: noContentResponse
//	404: errorResponse
//	501: errorResponse
func (t *Tasks) DeletTask(rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	t.l.Println("[DEBUG] deleting task with id ", id)

	err := t.db.Delete(data.TaskID(id))

	if err != nil {
		t.l.Printf("[Error] recored with specified id %s doesnt exist\n", id)
		rw.WriteHeader(http.StatusNotFound)
		x := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		buff, _ := json.Marshal(x)
		rw.Write(buff)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
