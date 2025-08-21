package utils

import "go.mongodb.org/mongo-driver/v2/bson"

func ValidadIdMongo(id string) (*bson.ObjectID, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectID, nil
}
