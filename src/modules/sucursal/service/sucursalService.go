package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/sucursal/dto"
	"comercial-backend/src/modules/sucursal/model"
	"comercial-backend/src/modules/sucursal/repository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegistrarSucursalService(body *dto.SucursalDto,ctx context.Context )error {

	var sucursal model.SucursalModel= model.SucursalModel {
		Nombre: body.Nombre,
		Direccion: body.Direccion,
		Fecha: time.Now(),
		Flag: enum.EstadoNuevo,
		
	}
	err := repository.CrearSucursalRepository(&sucursal, ctx)
	if(err != nil){
		return err
	}
	return nil

}

func ListarSucursalService(ctx context.Context )(*[]bson.M, error) {
	resultado, err := repository.ListarSucursalRepository(ctx)
	if(err != nil){
		return &[]bson.M{}, err
	}
	return resultado, nil

}