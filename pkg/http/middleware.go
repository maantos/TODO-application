package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maantos/todoApplication/pkg/domain"
)

func (h *Handler) ValidateBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		task := &domain.Task{}
		err := json.NewDecoder(r.Body).Decode(task)
		if err != nil {
			resp := Response{
				Code: http.StatusBadRequest,
				Msg:  fmt.Sprintf("invalid request body. %s", err.Error()),
			}
			respond(rw, r, &resp)
			return
		}

		ctx := context.WithValue(r.Context(), ContextBodyKey{}, task)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
