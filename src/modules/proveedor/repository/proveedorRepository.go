package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/structCore"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/proveedor/model"

	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CrearProveedorRepository(data *model.ProveedorModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Proveedor)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func ListarProveedorRepository(ci string, nombre string, celular string, empresa string, pagina int, limite int, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	collection := config.MongoDatabase.Collection(enum.Proveedor)
	filter := bson.M{"flag": enum.EstadoNuevo}
	if ci != "" {
		filter["ci"] = bson.D{
			{Key: "$regex", Value: ci},
			{Key: "$options", Value: "i"},
		}
	}
	if nombre != "" {
		filter["nombre"] = bson.D{
			{Key: "$regex", Value: nombre},
			{Key: "$options", Value: "i"},
		}
	}
	if celular != "" {
		filter["celular"] = bson.D{
			{Key: "$regex", Value: celular},
			{Key: "$options", Value: "i"},
		}
	}
	if empresa != "" {
		filter["empresa"] = bson.D{
			{Key: "$regex", Value: empresa},
			{Key: "$options", Value: "i"},
		}
	}
	opts := options.Find()
	opts.SetSkip(int64(utils.Skip(pagina, limite)))
	opts.SetLimit(int64(limite))
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	cantidad, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var proveedores []bson.M
	err = cursor.All(ctx, &proveedores)
	if err != nil {
		return nil, err
	}
	var paginas int = utils.CalcularPaginas(int(cantidad), limite)

	var resultado structCore.ResultadoPaginado = structCore.ResultadoPaginado{
		Data:    proveedores,
		Paginas: paginas,
	}
	return &resultado, nil
}
