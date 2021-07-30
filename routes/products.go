package routes

import (
	"context"
	"encoding/json"
	"fatihy101/filter_products/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	var products []*db.Product
	collection := getDB(r).ProductCollection()
	filters := getFilters(r)
	cursor, err := collection.Find(r.Context(), filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for cursor.Next(context.TODO()) {
		product := db.Product{}
		if err := cursor.Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	cursor.Close(context.TODO())

	json.NewEncoder(w).Encode(products)
}

func getFilters(r *http.Request) bson.D {
	var filters bson.D
	vars := r.URL.Query()
	fields := [6]string{"productName", "productCode", "productCompany", "productCategory", "productColor", "gender"}

	for _, fieldName := range fields {
		if fieldValue := vars.Get(fieldName); fieldValue != "" {
			filters = append(filters, createRegexFilter(fieldName, fieldValue))
		}
	}

	return filters
}

func createRegexFilter(key, pattern string) primitive.E {
	return primitive.E{Key: key, Value: primitive.Regex{Pattern: pattern, Options: "i"}}
}
