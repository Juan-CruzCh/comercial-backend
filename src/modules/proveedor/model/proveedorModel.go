package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProveedorModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty"` 
	CI        string        `bson:"ci"`            
	Nombre    string        `bson:"nombre"`        
	Apellidos string        `bson:"apellidos"`     
	Empresa   string        `bson:"empresa"`       
	Flag      string        `bson:"flag"`          
	Fecha     time.Time     `bson:"fecha"`         
}