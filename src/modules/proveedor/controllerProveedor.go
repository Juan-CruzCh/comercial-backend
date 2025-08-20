package proveedor

import (
	"comercial-backend/src/modules/proveedor/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func registrarProveedorController(c *gin.Context) {
	var proveedorDto dto.ProveedorDto
	err := c.ShouldBindJSON(&proveedorDto)
	if(err != nil) {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	err = registrarProveedorService(&proveedorDto, ctx)
	if(err != nil) {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	 c.JSON(http.StatusCreated, gin.H{"stattus":http.StatusCreated})
	 
}