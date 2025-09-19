package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type LogModel struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Usuario     bson.ObjectID `bson:"usuario"`
	Accion      string        `bson:"accion"`
	model       string        `bson:"model"`
	Ip          string        `bson:"ip"`
	Descripcion string        `bson:"descripcion"`
	Fecha       time.Time     `bson:"fecha"`
	Flag        enum.Estado   `bson:"flag"`
}
