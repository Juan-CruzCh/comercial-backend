package controller

import (
	"comercial-backend/src/modules/usuario/dto"
	"comercial-backend/src/modules/usuario/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CrearUsuarioController(c *gin.Context) {
	validate := validator.New()
	var body dto.UsuarioDto
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 20*time.Second)
	defer cancel()

	err = service.CrearUsuarioService(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})
}

func ListarUsuarioController(c *gin.Context) {

}

func ObtenerUsuarioController(c *gin.Context) {

}

func ActualizarUsuarioController(c *gin.Context) {

}

func EliminarUsuarioController(c *gin.Context) {

}
