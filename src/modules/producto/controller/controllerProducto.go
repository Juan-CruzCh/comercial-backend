package controller

import (
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegitrarProductoController(c *gin.Context) {
	validate := validator.New()
	var productoDto dto.ProductoDto
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := c.ShouldBindJSON(&productoDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(productoDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	categoriaID, err := bson.ObjectIDFromHex(productoDto.Categoria)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
		return
	}

	err = service.RegistrarProductoService(&productoDto, categoriaID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Producto registrado correctamente"})

}

func ListarProductoController(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	data, err := service.ListarProductoService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, data)
}
