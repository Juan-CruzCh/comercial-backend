package proveedor

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/proveedor/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func registrarProveedorController(c *gin.Context) {
	var proveedorDto dto.ProveedorDto
	err := c.ShouldBindJSON(&proveedorDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	err = registrarProveedorService(&proveedorDto, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})

}

func listarProveedorController(c *gin.Context) {
	pagina, limite, err := utils.Paginador(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	nombre := c.Query("nombre")
	ci := c.Query("ci")
	celular := c.Query("celular")
	empresa := c.Query("empresa")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	data, err := listarProveedorService(ci, nombre, celular, empresa, pagina, limite, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)

}
