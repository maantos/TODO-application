package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

// swagger:route POST /tasks tasks createTask
// Create a new TODO task
//
// responses:
//
//	200: createdTask
//	400: errorResponse
//	409: errorResponse
func (t *Tasks) CreateTask(rw http.ResponseWriter, r *http.Request) {

	task := r.Context().Value(KeyProduct{}).(*data.Task)
	err := t.db.Create(task)
	if err != nil {
		t.l.Printf("[Error] Error while creting task")
		rw.WriteHeader(http.StatusConflict)
		x := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		buff, _ := json.Marshal(x)
		rw.Write(buff)
		return
	}
	rw.WriteHeader(http.StatusOK)

	var res = struct {
		Id string `json:"id"`
	}{Id: string(task.ID)}

	data.ToJSON(res, rw)
	if err != nil {
		//if we get an error log it.
		t.l.Println("[ERROR] serializing product", err)
	}
}
