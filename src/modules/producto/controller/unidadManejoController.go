package controller

import (
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CrearUnidadController(c *gin.Context) {
	validate := validator.New()
	var body dto.UnidadManejoDto
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = service.CrearUnidadManejoService(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})
}

func ListarUnidadesController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	resultado, err := service.ListarUnidadManejoService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func ObtenerUnidadController(c *gin.Context) {

}

func ActualizarUnidadController(c *gin.Context) {

}

func EliminarUnidadController(c *gin.Context) {
}
