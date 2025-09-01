package middleware

import (
	"comercial-backend/src/modules/autenticacion/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidarTokenAtenticacion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("ctx")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": http.StatusForbidden})
			return
		}
		claims, err := utils.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": http.StatusForbidden})
			return
		}
		data := claims.(jwt.MapClaims)
		ctx.Set("usuario", data["usuario"])
		ctx.Set("sucursal", data["sucursal"])
		ctx.Next()
	}
}
