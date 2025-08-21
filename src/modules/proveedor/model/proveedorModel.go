package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProveedorModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty"` 
	CI        string        `bson:"ci"`            
	Nombre    string        `bson:"nombre"`  
	Celular    string        `bson:"celular"`        
	Apellidos string        `bson:"apellidos"`     
	Empresa   string        `bson:"empresa"`       
	Flag     enum.Estado        `bson:"flag"`          
	Fecha     time.Time     `bson:"fecha"`         
}