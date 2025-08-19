package producto

import (
	"comercial-backend/src/modules/producto/dto"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegitrarProductoController(c *gin.Context) {
	var productoDto dto.ProductoDto
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := c.ShouldBind(&productoDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	categoriaID, err := bson.ObjectIDFromHex(productoDto.Categoria)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría inválido"})
		return
	}
	fmt.Printf("Tipo de categoriaID: %T\n", categoriaID)
	err = RegistrarProductoService(&productoDto, categoriaID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Producto registrado correctamente"})

}

func ListarProductoController(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	data, err := ListarProductoService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, data)
}
