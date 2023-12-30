package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id           primitive.ObjectID `bson:"_id"`
	Heading      string             `json:"heading"`
	Price        string             `json:"price"`
	ProductMedia []ProductMedia     `json:"media"`
	Description  string             `json:"description"`
	Offers       []Offer            `json:"offers"`
}

type ProductMedia struct {
	MediaLinks string `json:""links`
}

type Offer struct {
	Offer string `json:"offer"`
}
