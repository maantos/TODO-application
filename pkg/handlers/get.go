package handlers

import (
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

// swagger:route GET /tasks tasks listTasks
// Get TODO tasks from database
// responses:
//	200: tasksResponse

func (t *Tasks) ListAll(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("[DEBUG] get all tasks")

	prods := t.db.List()
	if len(prods) < 1 {
		t.l.Printf("[DEBUG] list is empty")
		rw.WriteHeader(http.StatusOK)
		return
	}
	err := data.ToJSON(prods, rw)
	if err != nil {
		//if we get an error log it.
		t.l.Println("[ERROR] serializing product", err)
	}
}
