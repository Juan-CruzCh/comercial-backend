package controller

import (
	"comercial-backend/src/modules/autenticacion/dto"
	"comercial-backend/src/modules/autenticacion/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AutenticacionController(c *gin.Context) {
	var body dto.AutenticacionDto
	validate := validator.New()
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	token, err := service.AutenticacionService(&body, ctx)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(
		"ctx",
		token,
		4*60*60,
		"/",
		"",
		true,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}
