package handlers

import (
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

func (t *Tasks) ListAll(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("[DEBUG] get all tasks")

	prods := t.db.List()
	err := data.ToJSON(prods, rw)
	if err != nil {
		//if we get an error log it.
		t.l.Println("[ERROR] serializing product", err)
	}
}
