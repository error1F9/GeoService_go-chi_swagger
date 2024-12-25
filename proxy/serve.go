package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func (g *GeoService) Serve(address string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(NewReverseProxy("hugo", "1313").ReverseProxy)

	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from API"))
		})

		r.Post("/address/search", g.handleAddressSearch)
		r.Post("/address/geocode", g.handleAddressGeocode)

	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	return http.ListenAndServe(address, r)
}
