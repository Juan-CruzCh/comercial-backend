package controller

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/stock/dto"
	"comercial-backend/src/modules/stock/service"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegitrarStockController(c *gin.Context) {
	usuarioID, _, err := utils.Request(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	var body dto.IngresoStockData
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
	buscador := map[string]interface{}{}
	if codigo != "" {
		buscador["codigo"] = bson.Regex{Pattern: codigo, Options: "i"}
	}
	if nombreProducto != "" {
		buscador["nombre"] = bson.Regex{Pattern: nombreProducto, Options: "i"}
	}
	if categoria != "" {
		buscador["categoria"] = categoria
	}
	if unidadManejo != "" {
		buscador["unidadManejo"] = unidadManejo
	}

	fmt.Println("pagina", pagina, "limite", limite, "codig", buscador)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 20*time.Second)
	defer cancel()
	resultado, err := service.ListarStockService(ctx)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultado)
}
