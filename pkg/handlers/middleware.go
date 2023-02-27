package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/maantos/todoApplication/pkg/data"
)

func (t *Tasks) ValidateTask(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		task := &data.Task{}
		err := data.FromJSON(task, r.Body)
		if err != nil {
			t.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			x := struct {
				Message string `json:"message"`
			}{
				Message: err.Error(),
			}
			buff, _ := json.Marshal(x)
			rw.Write(buff)
			return
		}

		ctx := context.WithValue(r.Context(), TaskKey{}, task)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
