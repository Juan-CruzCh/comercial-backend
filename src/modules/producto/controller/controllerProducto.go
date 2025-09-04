package controller

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/service"
	"comercial-backend/src/modules/producto/structs"
	"context"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	categoriaID, err := utils.ValidadIdMongo(productoDto.Categoria)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
		return
	}
	unidadManejoID, err := utils.ValidadIdMongo(productoDto.UnidadManejo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
		return
	}
	producto, err := service.RegistrarProductoService(&productoDto, categoriaID, unidadManejoID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, producto)

}

func ListarProductoController(c *gin.Context) {
	pagina, limite, err := utils.Paginador(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	codigo := c.Query("codigo")
	nombreProducto := c.Query("nombreProducto")
	categoria := c.Query("categoria")
	unidadManejo := c.Query("unidadManejo")
	var buscador structs.FiltrosProductoStruct = structs.FiltrosProductoStruct{
		Codigo:         codigo,
		ProductoNombre: nombreProducto,
		Categoria:      categoria,
		UnidadManejo:   unidadManejo,
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	data, err := service.ListarProductoService(&buscador, pagina, limite, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
