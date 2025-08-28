package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UsuarioModel struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	CI       string        `bson:"ci"`
	Nombre   string        `bson:"nombre"`
	Apellidos string        `bson:"apellidos"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	Sucursal bson.ObjectID `bson:"sucursal"`
	Rol      string        `bson:"rol"`
	Fecha    time.Time     `bson:"fecha"`
	Flag     enum.Estado          `bson:"flag"`
}
