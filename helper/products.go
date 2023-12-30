package helper

import (
	"go.mongodb.org/mongo-driver/bson"
)

func FindDocument(columnName string, field interface{}) bson.D {
	return bson.D{{columnName, field}}
}

func FindProductsByText(text string) bson.D {
	return bson.D{{"$text", bson.D{{"$search", text}}}}
}
