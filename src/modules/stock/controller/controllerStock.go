package controller

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/stock/dto"
	"comercial-backend/src/modules/stock/service"
	structstock "comercial-backend/src/modules/stock/structStock"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegitrarStockController(c *gin.Context) {
	usuarioID, _, err := utils.Request(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	var body dto.IngresoStockDto
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	idIgreso, err := service.RegitrarStockService(&body, ctx, usuarioID)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ingreso": idIgreso})
}

func ListarStockController(c *gin.Context) {
	pagina, limite, err := utils.Paginador(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	codigo := c.Query("codigo")
	nombreProducto := c.Query("nombreProducto")
	categoria := c.Query("categoria")
	unidadManejo := c.Query("unidadManejo")
	var buscador structstock.FiltrosStock = structstock.FiltrosStock{
		Codigo:         codigo,
		ProductoNombre: nombreProducto,
		Categoria:      categoria,
		UnidadManejo:   unidadManejo,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 20*time.Second)
	defer cancel()
	resultado, err := service.ListarStockService(&buscador, pagina, limite, ctx)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultado)
}
