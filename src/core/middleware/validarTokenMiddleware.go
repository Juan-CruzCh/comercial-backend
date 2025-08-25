package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidarTokenAtenticacion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token bool = true
		if token {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"mensaje": "recurso proivido"})

	}
}
