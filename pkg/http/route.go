package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	cmiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-openapi/runtime/middleware"
)

func Routes(r chi.Router, h *Handler) {

	r.Use(cmiddleware.Logger)
	SwaggerFileServer(r, "/docs")
	r.Get("/tasks", h.GetTasks)
	r.Get("/", h.healthCheck)
	r.Delete("/tasks/{id:[0-9]+}", h.DeleteTask)

	// Private Routes
	// Require Authentication
	r.Group(func(r chi.Router) {
		r.Use(h.ValidateBody)
		r.Post("/tasks", h.CreateTask)
	})

}

func SwaggerFileServer(r chi.Router, path string) {
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle(path, sh)
	r.Handle("/swagger.json", http.FileServer(http.Dir("./docs")))
}
