package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
)

func Routes() http.Handler {

	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)

	SwaggerFileServer(r, "/docs")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/tasks", a.th.ListAll)
	r.Delete("/tasks/{id:[0-9]+}", a.th.DeletTask)

	// Private Routes
	// Require Authentication
	r.Group(func(r chi.Router) {
		r.Use(a.th.ValidateTask)
		r.Post("/tasks", a.th.CreateTask)
	})

	return r
}

func SwaggerFileServer(r chi.Router, path string) {
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle(path, sh)
	r.Handle("/swagger.json", http.FileServer(http.Dir("./")))
}
