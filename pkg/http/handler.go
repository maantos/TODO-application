package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maantos/todoApplication/pkg/domain"
)

type Handler struct {
	tskSvc domain.TaskService
}

func NewHandler(svc domain.TaskService) *Handler {
	return &Handler{
		tskSvc: svc,
	}
}

// swagger:route GET /tasks tasks
// Get TODO tasks from database
// responses:
//
//	200: tasksResponse
func (h *Handler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	// t.l.Println("[DEBUG] get all tasks")

	prods := h.tskSvc.List()
	if len(prods) < 1 {
		rw.WriteHeader(http.StatusOK)
		return
	}
	e := json.NewEncoder(rw)
	err := e.Encode(prods)
	if err != nil {
		//if we get an error log it.
		fmt.Println(fmt.Errorf("marshalling response to json failed, err: %v", err))
	}
}

// swagger:route POST /tasks tasks createTask
// Create a new TODO task
//
// responses:
//
//	200: createdTask
//	400: errorResponse
//	409: errorResponse
func (h *Handler) CreateTask(rw http.ResponseWriter, r *http.Request) {

	task := r.Context().Value(ContextBodyKey{}).(*domain.Task)
	err := h.tskSvc.Create(task)
	if err != nil {
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

	e := json.NewEncoder(rw)
	err = e.Encode(res)
	if err != nil {
		//if we get an error log it.
		fmt.Println(fmt.Errorf("marshalling response to json failed, err: %v", err))
	}
}

// swagger:route DELETE /tasks/{id} tasks deleteTask
// Removes task from the database
//
// responses:
//
//	204: noContentResponse
//	404: errorResponse
func (h *Handler) DeleteTask(rw http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := h.tskSvc.Delete(domain.TaskID(id))

	if err != nil {
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

// swagger:route GET /
// healthcheck
//
// responses:
//
//	200: StatusOK
func (h *Handler) healthCheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Site is up."))
}
