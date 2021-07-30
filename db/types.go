package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBHandle struct {
	mdb *mongo.Database
}

type Product struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductCode     string             `bson:"ProductCode" json:"ProductCode"`
	ProductName     string             `bson:"productName" json:"productName"`
	Price           float64            `bson:"priceSale" json:"priceSale"`
	ImageLink       string             `bson:"imageLink" json:"imageLink"`
	ProductCompany  string             `bson:"productCompany" json:"productCompany"`
	ProductCategory string             `bson:"productCategory" json:"productCategory"`
	ProductColor    string             `bson:"productColor" json:"productColor"`
	Gender          string             `bson:"gender" json:"gender"`
	Website         string             `bson:"webSite" json:"webSite"`
}
