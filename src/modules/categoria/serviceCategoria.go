package categoria

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/categoria/dto"
	"comercial-backend/src/modules/categoria/model"
	"comercial-backend/src/modules/categoria/repository"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func crearCategoriaService(categoria *dto.CategoriaDto, ctx context.Context) error {

	data := model.Categoria{
		Nombre: categoria.Nombre,
		Fecha:  time.Now(),
		Flag:   enum.EstadoNuevo,
	}
	err := repository.CrearCategoriaRepository(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}
func ListarCategoriaService(ctx context.Context) (*[]bson.M, error) {
	data, err := repository.ListarCategoriaRepository(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func ObtenerCategoriaService(c *gin.Context) {

}

func ActualizarCategoriaService(c *gin.Context) {

}

func eliminarCategoriaService(categoriID *bson.ObjectID, ctx context.Context) error {
	err := repository.EliminarCategoriaRepository(categoriID, ctx)
	if err != nil {
		return err
	}
	return nil
}
