package main

import (
	"log"
	"net/http"

	"github.com/cloudykit/jet"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(a.LoadSession)

	if a.debug {
		mux.Use(middleware.Logger)
	}
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		a.session.Put(r.Context(), "test", "titas")
		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		vars := make(jet.VarMap)
		vars.Set("test", a.session.GetString(r.Context(), "test"))
		err := a.render(w, r, "index", vars)
		if err != nil {
			log.Fatalln(err)
		}
	})
	return mux
}
