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
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := service.ObtenerUsuarioIdService(ID, ctx)
	c.JSON(http.StatusOK, gin.H{"status": data})
}

func ActualizarUsuarioController(c *gin.Context) {
	validate := validator.New()
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var body dto.UsuarioDto
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.ActualizarUsuarioService(ID, &body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	id, existe := c.Get("usuario")

	if !existe {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al convertir el ID a string"})
		return
	}
	ID, err := utils.ValidadIdMongo(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error  el ID a Mongo"})
	}
	data, err := service.ObtenerUsuarioIdService(ID, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error  al obtener el usuario"})
	}
	c.JSON(http.StatusOK, data)

}
