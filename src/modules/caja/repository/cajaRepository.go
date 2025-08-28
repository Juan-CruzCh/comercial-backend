package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/caja/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func AbrirCajaRepository(data *model.CajaModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func VerificarCajaAbierto(usuario *bson.ObjectID, ctx context.Context) error {
	//collection := config.MongoDatabase.Collection(enum.Caja)
	hoy := time.Now()
	fechaInicio := time.Date(hoy.Year(), hoy.Month(), hoy.Day(), 0, 0, 0, 0, hoy.Location())
	fechaFin := time.Date(hoy.Year(), hoy.Month(), hoy.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), hoy.Location())
	fmt.Println(fechaFin)
	fmt.Println(fechaInicio)
	/*filter := bson.M{
		"usuario": usuario,
		"fechaApertura": bson.M{
			"$gte": "",
			"$lt":  "",
		},
		"estado": enum.Abierto,
	}
	cantidad, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Println(cantidad)
	if cantidad > 0 {
		return errors.New("La caja ya se encuentra abierta")
	}*/
	return nil
}
