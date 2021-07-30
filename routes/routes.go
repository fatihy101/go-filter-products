package routes

import (
	"encoding/json"
	"fatihy101/filter_products/db"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(db *db.DBHandle) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(AllowOrigin)
	router.Use(DBMiddleware(db))
	router.Use(JSONResponseMiddleware)
	router.Route("/", allRoutes)

	return router
}

func allRoutes(r chi.Router) {
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(map[string]string{"server": "active"})
	})
	r.Route("/products", productsRoutes)
}

func productsRoutes(r chi.Router) {
	r.Get("/", getProducts)
}
