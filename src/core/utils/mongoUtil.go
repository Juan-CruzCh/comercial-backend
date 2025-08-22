package utils

import "go.mongodb.org/mongo-driver/v2/bson"

func Lookup(from, localField, foreignField, as string) bson.D {
	var pipelineMongo = bson.D{
		{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: from},
				{Key: "localField", Value: localField},
				{Key: "foreignField", Value: foreignField},
				{Key: "as", Value: as},
			},
		},
	}
	return pipelineMongo
}

func Unwind(path string, preserveNullAndEmptyArrays bool) bson.D {
	return bson.D{
		{
			Key: "$unwind",
			Value: bson.D{
				{Key: "path", Value: path},
				{Key: "preserveNullAndEmptyArrays", Value: preserveNullAndEmptyArrays},
			},
		},
	}

}

func ArrayElemAt(arrayElemAtStage string, indice int) bson.D {
	var pipeline bson.D = bson.D{
		{Key: "$arrayElemAt", Value: bson.A{arrayElemAtStage, indice}},
	}
	return pipeline
}
