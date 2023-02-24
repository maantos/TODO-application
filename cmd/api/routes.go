package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *application) routes() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
		r.Get("/tasks", a.th.ListAll)
		r.Delete("/tasks/{id:[0-9]+}", a.th.DeletTask)
	})

	// Private Routes
	// Require Authentication
	r.Group(func(r chi.Router) {
		r.Use(a.th.ValidateTask)
		r.Post("/tasks", a.th.CreateTask)
	})

	return r
}
