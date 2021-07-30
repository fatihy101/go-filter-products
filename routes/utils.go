package routes

import (
	"fatihy101/filter_products/db"
	"net/http"
)

func getDB(r *http.Request) *db.DBHandle {
	return r.Context().Value(DBContext).(*db.DBHandle)
}
