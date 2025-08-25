package utils

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ValidadIdMongo(id string) (*bson.ObjectID, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID de mongo invalido")
	}
	return &objectID, nil
}
