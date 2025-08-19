package categoria

import (
	"comercial-backend/src/modules/categoria/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CrearCategoria(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var categoria dto.CategoriaDto
	err := c.ShouldBind(&categoria)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = crearCategoriaService(&categoria, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensage": "registrado"})

}

func ListarCategoria(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	data, err := ListarCategoriaService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, data)

}
