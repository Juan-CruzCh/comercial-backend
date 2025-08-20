package main

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/modules/almacen"
	"comercial-backend/src/modules/categoria"
	"comercial-backend/src/modules/producto"

	"comercial-backend/src/modules/proveedor"
	"comercial-backend/src/modules/stock"
	"comercial-backend/src/modules/usuario"
	"comercial-backend/src/modules/venta"

	"github.com/gin-gonic/gin"
)

func main() {

	var url string = "mongodb://kanna:kanna@localhost:27017/comision?authSource=admin"
	//var url string = "mongodb://localhost:27017"
	config.ConnectMongo(url, "ventas")

	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := config.MongoClient.Disconnect(ctx)

	if err != nil {
		fmt.Println("⚠️ Error al desconectar Mongo:", err)
	} else {
		fmt.Println("👋 Desconectado de MongoDB")
	}*/

	router := gin.Default()
	api := router.Group("api")
	// router categorias
	categoria.RouterCategoria(api)
	proveedor.RouterProveedor(api)
	stock.RouterStock(api)
	producto.RouterProducto(api)
	usuario.UsuarioRouter(api)
	venta.VentaRouter(api)
	almacen.AlmacenRouter(api)
	router.Run(":5000")
}
