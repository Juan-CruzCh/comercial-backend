package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Request(c *gin.Context) (usuarioID *bson.ObjectID, sucursalID *bson.ObjectID, err error) {
	usuarioRaw, exists := c.Get("usuario")
	if !exists {
		return nil, nil, errors.New("no se encontró el valor 'usuario' en la solicitud")
	}

	sucursalRaw, exists := c.Get("sucursal")
	if !exists {
		return nil, nil, errors.New("no se encontró el valor 'sucursal' en la solicitud")
	}

	usuarioID, err = ValidadIdMongo(usuarioRaw.(string))
	if err != nil {
		return nil, nil, errors.New("error al validar el ID de usuario: " + err.Error())
	}

	sucursalID, err = ValidadIdMongo(sucursalRaw.(string))
	if err != nil {
		return nil, nil, errors.New("error al validar el ID de sucursal: " + err.Error())
	}

	return usuarioID, sucursalID, nil
}
