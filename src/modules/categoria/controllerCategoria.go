package categoria

import (
	"comercial-backend/src/modules/categoria/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)
	

func CrearCategoria(c *gin.Context) {
	var categoria dto.CategoriaDto
	err := c.ShouldBind(&categoria)
	 if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
 	_, err =crearCategoriaService(&categoria)
	 if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
	c.JSON(http.StatusCreated, gin.H{"mensage":"registrado"})	
	
}

func ListarCategoria(c *gin.Context) {
	
 	data, err :=ListarCategoriaService()
	 if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
	c.JSON(http.StatusCreated, gin.H{"data":data})	
	
}