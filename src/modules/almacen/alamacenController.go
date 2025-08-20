package almacen

import (
	"comercial-backend/src/modules/almacen/dto"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func registrarAlmacenController(c *gin.Context) {
	validate := validator.New()
	var body dto.AlmacenDto
	err := c.ShouldBindJSON(&body)
	body.Nombre = strings.ToUpper(body.Nombre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	err = registrarAlmacenService(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "registrado"})

}

func listarAlmacenesController(c *gin.Context) {

}

func obtenerAlmacenController(c *gin.Context) {

}

func actualizarAlmacenController(c *gin.Context) {

}

func eliminarAlmacenController(c *gin.Context) {

}
