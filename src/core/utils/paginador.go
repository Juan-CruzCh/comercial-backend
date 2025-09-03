package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Paginador(c *gin.Context) (pagina int, limite int, err error) {
	paginaStr := c.DefaultQuery("pagina", "1")
	limiteStr := c.DefaultQuery("limite", "20")

	pagina, err = strconv.Atoi(paginaStr)

	if err != nil {
		return 0, 0, errors.New("Ingrese el numero pagina")
	}
	limite, err = strconv.Atoi(limiteStr)
	if err != nil {
		return 0, 0, errors.New("Ingrese el numero limite")
	}
	return pagina, limite, nil

}
