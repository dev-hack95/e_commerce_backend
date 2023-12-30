package models

import (
	"context"
	"e_commerce_backend/helper"
	"e_commerce_backend/utilities"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	Id                   primitive.ObjectID `bson:"_id"`
	ProductName          string             `json:"product_name"`
	Price                string             `json:"price"`
	ProductMedia         string             `json:"media"`
	Description          string             `json:"description"`
	Offers               string             `json:"offers"`
	ProductNameEmbedding []float64          `json:"product_embedding"`
}

func AddProduct(u *Product) (*mongo.InsertOneResult, error) {
	client, err := utilities.GetDBInstance()
	if err != nil {
		panic(err)
	}
	collection := client.Database("products").Collection("products")
	id, err := collection.InsertOne(context.TODO(), &u)
	if err != nil {
		panic(err)
	}
	return id, nil
}

func GetProductById(productId string) (*Product, error) {
	client, err := utilities.GetDBInstance()
	if err != nil {
		return nil, err
	}
	collection := client.Database("products").Collection("products")
	filter := helper.FindDocument("_id", cast.ToString(productId))
	var result Product
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func SearchProducts(text string) ([]Product, error) {
	client, err := utilities.GetDBInstance()
	if err != nil {
		return nil, err
	}
	collection := client.Database("products").Collection("products")
	searchFilter := helper.FindProductsByText(text)
	cursor, err := collection.Find(context.TODO(), searchFilter)

	if err != nil {
		return nil, err
	}
	var results []Product

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
