package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(rol []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(rol)
	}
}
