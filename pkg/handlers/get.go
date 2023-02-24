package handlers

import (
	"fmt"
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

func (t *Tasks) ListAll(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("[DEBUG] get all tasks")

	prods := t.db.List()
	fmt.Println("herer")
	if len(prods) < 1 {
		t.l.Printf("[DEBUG] list is empty")
		rw.WriteHeader(http.StatusOK)
		return
	}
	fmt.Println("herereeee")
	err := data.ToJSON(prods, rw)
	if err != nil {
		//if we get an error log it.
		t.l.Println("[ERROR] serializing product", err)
	}
}
