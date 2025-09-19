package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/log/model"
	"context"
)

func RegistrarLogRespository(data model.LogModel, ctx context.Context) {
	collection := config.MongoDatabase.Collection(enum.Log)
	collection.InsertOne(ctx, data)

}
