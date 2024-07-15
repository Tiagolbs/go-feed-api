package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Get("/{id}", app.showPostHandler)
			r.Get("/", app.listPostsHandler)
			r.Put("/{id}", app.editPostHandler)
			r.Delete("/{id}", app.deletePostHandler)
		})
	})

	return router
}
