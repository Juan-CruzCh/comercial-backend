package controller

import (
	"comercial-backend/src/core/utils"
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	data, err := service.ListarUsuarioService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)

}

func ObtenerUsuarioController(c *gin.Context) {

}

func ActualizarUsuarioController(c *gin.Context) {

}

func EliminarUsuarioController(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.EliminarUsuarioService(ID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}
func CerrarAuntenticacionUsuarioController(c *gin.Context) {
	_, existe := c.Get("usuario")
	if !existe {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	c.SetCookie("ctx", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}
func VerificarAutenticacionUsuarioController(c *gin.Context) {
	_, existe := c.Get("usuario")

	if !existe {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}
